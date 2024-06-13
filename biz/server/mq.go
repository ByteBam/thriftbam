package server

import (
	"context"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/biz/middleware"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/ByteBam/thirftbam/pkg/utils/server/mq"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/bytedance/sonic"
	"github.com/spf13/viper"
)

func NewMQServer(
	logger *log.Logger,
	conf *viper.Viper,
	service service.AnalyzeService,
) *mq.Server {
	s := mq.NewServer(
		middleware.NewRMQ(conf),
		logger,
		service,
	)

	err := s.Subscribe(conf.GetString("rocketmq.topic"), consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			var req v1.AnalyzeRequest
			logger.Info(string(msg.Body))
			if err := sonic.Unmarshal(msg.Body, &req); err != nil {
				logger.Error(err.Error())
				return consumer.ConsumeRetryLater, err
			}
			url, err := s.Service.GetUrl(ctx, &req)
			if err != nil {
				logger.Error(err.Error())
				return consumer.Rollback, err
			}
			loading, err := s.Service.Loading(ctx, url)
			if err != nil {
				logger.Error(err.Error())
				return consumer.Rollback, err
			}
			nums, err := s.Service.Parser(ctx, *loading, req.BranchId)
			if err != nil {
				logger.Error(err.Error())
				return consumer.Rollback, err
			}
			if err = s.Service.InterfaceNums(ctx, req.BranchId, nums); err != nil {
				logger.Error(err.Error())
				return consumer.Rollback, err
			}
		}
		return consumer.ConsumeSuccess, nil
	},
	)
	if err != nil {
		return nil
	}

	return s
}
