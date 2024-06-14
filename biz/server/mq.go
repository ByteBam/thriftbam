package server

import (
	"github.com/ByteBam/thirftbam/biz/middleware"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/ByteBam/thirftbam/pkg/utils/server/mq"
	"github.com/spf13/viper"
)

func NewMQServer(
	logger *log.Logger,
	conf *viper.Viper,
	service service.AnalyzeService,
) *mq.Server {
	s := mq.NewServer(
		middleware.NewRMQConsumer(conf, logger, service),
		logger,
		service,
	)

	return s
}
