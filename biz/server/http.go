package server

import (
	"github.com/ByteBam/thirftbam/biz/handler"
	"github.com/ByteBam/thirftbam/util/log"
	"github.com/ByteBam/thirftbam/util/server/http"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	analyzeHandler *handler.AnalyzeHandler,
) *http.Server {
	s := http.NewServer(
		server.Default(server.WithHostPorts(conf.GetString("app.http.addr"))),
		logger,
	)

	s.Hertz.GET("/analyze", analyzeHandler.Analyze)
	return s
}
