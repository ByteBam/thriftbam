package server

import (
	"context"
	"fmt"
	"github.com/ByteBam/thirftbam/biz/handler"
	_ "github.com/ByteBam/thirftbam/docs"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/ByteBam/thirftbam/pkg/utils/server/http"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
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

	url := fmt.Sprintf("http://%s/swagger/doc.json", conf.GetString("app.http.addr"))
	s.Hertz.Any("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "pong")
	})
	s.Hertz.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, swagger.URL(url)))
	v1 := s.Hertz.Group("/api/v1")
	v1.POST("/analyze", analyzeHandler.Analyze)

	return s
}
