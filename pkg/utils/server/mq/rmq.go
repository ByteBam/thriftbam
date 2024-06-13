package mq

import (
	"context"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/apache/rocketmq-client-go/v2"
	"time"
)

type Server struct {
	rocketmq.PushConsumer
	logger  *log.Logger
	Service service.AnalyzeService
}

type Option func(s *Server)

func NewServer(c rocketmq.PushConsumer, logger *log.Logger, analyzeService service.AnalyzeService, opts ...Option) *Server {
	s := &Server{
		PushConsumer: c,
		logger:       logger,
		Service:      analyzeService,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Sugar().Info("Starting server...")

	err := s.PushConsumer.Start()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.PushConsumer.Shutdown(); err != nil {
		return err
	}

	s.logger.Sugar().Info("Server exiting")
	return nil
}
