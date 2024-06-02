package service

import (
	"context"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
)

type AnalyzeService interface {
	Download(ctx context.Context, request *v1.AnalyzeRequest) (string, error)
	AnalyzeFile(ctx context.Context, file string) (string, error)
	DeleteFile(ctx context.Context, file string) error
}

type analyzeService struct {
	*Service
}

func NewAnalyzeService(
	service *Service,
) AnalyzeService {
	return &analyzeService{
		Service: service,
	}
}

func (a analyzeService) Download(ctx context.Context, request *v1.AnalyzeRequest) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (a analyzeService) AnalyzeFile(ctx context.Context, file string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (a analyzeService) DeleteFile(ctx context.Context, file string) error {
	// TODO implement me
	panic("implement me")
}
