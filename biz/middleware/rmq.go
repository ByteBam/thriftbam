package middleware

import (
	"context"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/bytedance/sonic"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewRMQConsumer(
	conf *viper.Viper,
	logger *log.Logger,
	service service.AnalyzeService,
) []rocketmq.PushConsumer {
	var consumers []rocketmq.PushConsumer
	if err := rlog.SetOutputPath(conf.GetString("rocketmq.log.log_file_name")); err != nil {
		panic(err)
	}

	for _, consumerConf := range conf.Get("rocketmq.consumers").([]interface{}) {
		c, err := rocketmq.NewPushConsumer(
			consumer.WithNameServer([]string{conf.GetString("rocketmq.nameSrv")}),
			consumer.WithConsumerModel(consumer.Clustering),
			consumer.WithGroupName(consumerConf.(map[string]interface{})["group"].(string)),
		)

		if err != nil {
			panic(err)
		}

		if err = c.Subscribe(
			consumerConf.(map[string]interface{})["topic"].(string),
			consumer.MessageSelector{},
			func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				// traverse the messages in rocketmq
				for _, msg := range msgs {
					var req v1.AnalyzeRequest

					// unmarshal the message body
					if err := sonic.Unmarshal(msg.Body, &req); err != nil {
						logger.WithContext(ctx).Error("sonic.Unmarshal error", zap.Error(err))
						return consumer.ConsumeRetryLater, err
					}

					// get the idl file url for loading
					url, err := service.GetUrl(ctx, &req)
					if err != nil {
						logger.WithContext(ctx).Error("AnalyzeService.GetUrl error", zap.Error(err))
						return consumer.Rollback, err
					}

					// load the idl file
					loading, err := service.Loading(ctx, url)
					if err != nil {
						logger.WithContext(ctx).Error("AnalyzeService.Loading error", zap.Error(err))
						return consumer.Rollback, err
					}

					// parse the idl file
					nums, err := service.Parser(ctx, *loading, req.BranchId)
					if err != nil {
						// if the idl file not found, update status is failed and return the message consumption success
						if dbErr := service.ParserError(ctx, req.BranchId); dbErr != nil {
							return consumer.Rollback, dbErr
						}
						if err.Error() == "the idl file not found" {
							return consumer.ConsumeSuccess, nil
						}
						logger.WithContext(ctx).Error("AnalyzeService.Parser error", zap.Error(err))
						return consumer.Rollback, err
					}
					if err = service.InterfaceNums(ctx, req.BranchId, nums); err != nil {
						logger.WithContext(ctx).Error("AnalyzeService.InterfaceNums error", zap.Error(err))
						return consumer.Rollback, err
					}
				}
				return consumer.ConsumeSuccess, nil
			},
		); err != nil {
			logger.Error("NewMQServer.subscribe error", zap.Error(err))
			return nil
		}
		logger.Info("NewMQServer.subscribe success", zap.String("topic", consumerConf.(map[string]interface{})["topic"].(string)))
		consumers = append(consumers, c)
	}

	return consumers
}
