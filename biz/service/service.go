package service

import "github.com/ByteBam/thirftbam/util/log"

type Service struct {
	logger *log.Logger
}

func NewService(
	logger *log.Logger,
) *Service {
	return &Service{
		logger: logger,
	}
}
