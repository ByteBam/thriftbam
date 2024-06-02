//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/ByteBam/thirftbam/biz/app"
	"github.com/ByteBam/thirftbam/biz/handler"
	"github.com/ByteBam/thirftbam/biz/server"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/util/log"
	"github.com/ByteBam/thirftbam/util/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewAnalyzeService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAnalyzeHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

func newApp(
	httpServer *http.Server,
	conf *viper.Viper,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName(conf.GetString("app.name")),
	)
}

func NewWire(viper2 *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		serviceSet,
		handlerSet,
		serverSet,
		newApp,
	))
}
