package app

import (
	"context"
	"github.com/ByteBam/thirftbam/pkg/utils/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		go opt(a)
	}
	return a
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = append(a.servers, servers...)
	}
}

func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func (a *App) Run(ctx context.Context) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for _, srv := range a.servers {
		go func(srv server.Server) {
			err := srv.Start(ctx)
			if err != nil {
				log.Printf("Server start err: %v", err)
			}
		}(srv)
	}

	select {
	case <-signals:
		log.Println("Received termination signal")
	case <-ctx.Done():
		log.Println("Context done")
	}

	for _, srv := range a.servers {
		if err := srv.Stop(ctx); err != nil {
			log.Printf("Server stop err: %v", err)
		}
	}

	return nil
}
