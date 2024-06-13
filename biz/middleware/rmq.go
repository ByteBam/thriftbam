package middleware

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/spf13/viper"
)

func NewRMQ(conf *viper.Viper) rocketmq.PushConsumer {
	fmt.Printf("NewRMQ: %v\n", conf.GetString("rocketmq.nameSrv"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{conf.GetString("rocketmq.nameSrv")}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(conf.GetString("rocketmq.group")),
	)

	if err != nil {
		panic(err)
	}
	return c
}
