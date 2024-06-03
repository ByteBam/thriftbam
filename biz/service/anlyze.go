package service

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/util/args"
	"github.com/ByteBam/thirftbam/util/download"
	"github.com/ByteBam/thirftbam/util/parser"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"os"
	"sync"
)

type AnalyzeService interface {
	Download(ctx context.Context, request *v1.AnalyzeRequest) error
	Analyze(ctx context.Context, path string) (int, error)
	Delete(ctx context.Context, path string) error
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

	if err = download.File(ctx, contents, request.Owner); err != nil {
		return err
	}

	return nil
}

func (a analyzeService) Analyze(ctx context.Context, path string) (int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open directory:", err)
		return 0, err
	}
	defer dir.Close()

	// 读取目录中的文件列表
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Failed to read directory contents:", err)
		return 0, err
	}

	var wg sync.WaitGroup
	var once sync.Once
	result := make(chan *parser.Thrift, len(fileInfos))
	rate := make(chan struct{}, 10) // 控制并发数

	// 启动一个单独的 Goroutine 来处理结果
	var ast *parser.Thrift
	var parseErr error
	done := make(chan struct{})
	go func() {
		for res := range result {
			if ast == nil {
				ast = res
			}
		}
		close(done)
	}()

	for _, info := range fileInfos {
		if !info.IsDir() { // 确保只处理文件
			rate <- struct{}{}
			wg.Add(1)
			go func(info os.FileInfo) {
				defer func() {
					<-rate
					wg.Done()
				}()
				var arguments args.Arguments
				arguments.IDL = path + "/" + info.Name()
				parsedAst, err := parser.ParseFile(arguments.IDL, arguments.Includes, true)
				if err != nil {
					once.Do(func() {
						parseErr = err
						cancel()
					})
					return
				}
				result <- parsedAst
			}(info)
		}
	}

	wg.Wait()
	close(result)
	<-done

	if parseErr != nil {
		return 0, parseErr
	}
	if ast == nil {
		return 0, fmt.Errorf("no services found")
	}
	// TODO The parameters of the build request and the corresponding parameters are stored in the database
	var count int
	for _, service := range ast.Services {
		count = count + len(service.Functions)
	}
	return count, nil
}

func (a analyzeService) Delete(ctx context.Context, path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
