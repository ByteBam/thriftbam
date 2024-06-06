// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func NewWire(viper2 *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	handlerHandler := handler.NewHandler(logger)
	sidSid := sid.NewSid()
	db := repository.NewDB(viper2, logger)
	client := repository.NewRedis(viper2)
	repositoryRepository := repository.NewRepository(logger, db, client)
	transaction := repository.NewTransaction(repositoryRepository)
	serviceService := service.NewService(logger, sidSid, transaction)
	queryRepository := repository.NewQueryRepository(repositoryRepository)
	captchaRepository := repository.NewCaptchaRepository(repositoryRepository)
	analyzeService := service.NewAnalyzeService(serviceService, queryRepository, captchaRepository)
	analyzeHandler := handler.NewAnalyzeHandler(handlerHandler, analyzeService)
	httpServer := server.NewHTTPServer(logger, viper2, analyzeHandler)
	appApp := newApp(httpServer, viper2)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewTransaction, repository.NewQueryRepository, repository.NewCaptchaRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewAnalyzeService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewAnalyzeHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

func newApp(
	httpServer *http.Server,
	conf *viper.Viper,
) *app.App {
	return app.NewApp(app.WithServer(httpServer), app.WithName(conf.GetString("app.name")))
}