package http

import (
	"context"
	"github.com/ByteBam/thirftbam/util/log"
	"github.com/cloudwego/hertz/pkg/app/server"
	"time"
)

type Server struct {
	*server.Hertz
	logger *log.Logger
}

type Option func(s *Server)

func NewServer(hertz *server.Hertz, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Hertz:  hertz,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	s.Hertz.Spin()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.SetCustomSignalWaiter(func(err chan error) error {
		return nil
	})

	s.logger.Sugar().Info("Server exiting")
	return nil
}
