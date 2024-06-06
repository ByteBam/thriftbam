//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/ByteBam/thirftbam/biz/app"
	"github.com/ByteBam/thirftbam/biz/handler"
	"github.com/ByteBam/thirftbam/biz/repository"
	"github.com/ByteBam/thirftbam/biz/server"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/ByteBam/thirftbam/pkg/utils/server/http"
	"github.com/ByteBam/thirftbam/pkg/utils/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewQueryRepository,
	repository.NewCaptchaRepository,
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

func NewWire(viper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		newApp,
	))
}
