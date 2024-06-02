package service

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/util/download"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

type AnalyzeService interface {
	Download(ctx context.Context, request *v1.AnalyzeRequest) error
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

func (a analyzeService) Download(ctx context.Context, request *v1.AnalyzeRequest) error {
	// TODO get ack by userID
	// 拼接url
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s?access_token=%s&ref=%s", request.Owner, request.Repo, request.Path, request.AccessToken, request.Ref)
	// 发起 GET 请求
	reps, err := http.Get(url)
	if err != nil {
		return err
	}
	if reps.StatusCode != http.StatusOK {
		switch reps.StatusCode {
		case http.StatusUnauthorized:
			return errors.New("failed to send request: unauthorized")
		}
	}

	body, err := io.ReadAll(reps.Body)
	if err != nil {
		return err
	}

	var contents []download.Content
	if err = sonic.Unmarshal(body, &contents); err != nil {
		return err
	}
	defer reps.Body.Close()

	if err = download.File(contents, request.Owner); err != nil {
		return err
	}

	return nil
}

func (a analyzeService) AnalyzeFile(ctx context.Context, file string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (a analyzeService) DeleteFile(ctx context.Context, file string) error {
	// TODO implement me
	panic("implement me")
}
