package service

import (
	"github.com/ByteBam/thirftbam/util/log"
	"github.com/ByteBam/thirftbam/util/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
}

func NewService(
	logger *log.Logger,
	sid *sid.Sid,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
	}
}
