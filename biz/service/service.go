package service

import (
	"github.com/ByteBam/thirftbam/biz/repository"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"github.com/ByteBam/thirftbam/pkg/utils/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	tm     repository.Transaction
}

func NewService(
	logger *log.Logger,
	sid *sid.Sid,
	tm repository.Transaction,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
