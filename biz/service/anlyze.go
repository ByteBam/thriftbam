package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/biz/model"
	"github.com/ByteBam/thirftbam/biz/repository"
	"github.com/ByteBam/thirftbam/pkg/parser"
	"github.com/ByteBam/thirftbam/pkg/util/download"
	"github.com/ByteBam/thirftbam/pkg/util/interface_info"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"sync"
)

var USER_ACCUSER_ACCESSION_KEY = "USER_ACCUSER_ACCESSION_KEY"

type AnalyzeService interface {
	GetUrl(ctx context.Context, request *v1.AnalyzeRequest) (string, error)
	Loading(ctx context.Context, url string) (*[]download.IDL, error)
	Parser(ctx context.Context, idls []download.IDL, branchId string) (int, error)
}

type analyzeService struct {
	*Service
	db  repository.QueryRepository
	rdb repository.CaptchaRepository
}

func (a *analyzeService) GetUrl(ctx context.Context, request *v1.AnalyzeRequest) (string, error) {
	key := fmt.Sprintf("%s%s", USER_ACCUSER_ACCESSION_KEY, request.UserId)
	ack, err := a.rdb.GetTokenByRDS(ctx, key)
	if err != nil {
		return "", err
	}
	ref, err := a.db.GetBranchByID(ctx, request.BranchId)
	if err != nil {
		return "", err
	}
	// 拼接url
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s?access_token=%s&ref=%s", request.Owner, request.Repo, request.Path, ack, ref)
	return url, nil
}

func (a *analyzeService) Loading(ctx context.Context, url string) (*[]download.IDL, error) {
	// 发起 GET 请求
	reps, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if reps.StatusCode != http.StatusOK {
		switch reps.StatusCode {
		case http.StatusUnauthorized:
			return nil, errors.New("failed to send request: unauthorized")
		}
	}

	body, err := io.ReadAll(reps.Body)
	if err != nil {
		return nil, err
	}

	var contents []download.Content
	if err = sonic.Unmarshal(body, &contents); err != nil {
		return nil, err
	}
	defer reps.Body.Close()

	idls, err := download.FileContent(ctx, &contents)
	if err != nil {
		return nil, err
	}

	return idls, nil
}

func (a *analyzeService) Parser(ctx context.Context, idls []download.IDL, branchId string) (int, error) {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	var wg sync.WaitGroup
	var once sync.Once
	var err error
	var structure sync.Map
	result := make(chan *parser.Thrift, len(idls))
	rate := make(chan struct{}, 10) // 控制并发数

	// 启动一个单独的 Goroutine 来处理结果
	var ast *parser.Thrift
	done := make(chan struct{})
	go func() {
		for res := range result {
			if ast == nil {
				ast = res
			}
		}
		close(done)
	}()

	for _, idl := range idls {
		if len(idl.Content) > 0 { // 确保文件非空
			rate <- struct{}{}
			wg.Add(1)
			go func(idl download.IDL) {
				defer func() {
					<-rate
					wg.Done()
				}()
				parsedAst, e := parser.ParseString(idl.Path, idl.Content)
				if e != nil {
					once.Do(func() {
						err = e
						cancel(e)
					})
					return
				}
				for _, s := range parsedAst.Structs {
					structure.Store(s.Name, s)
				}
				result <- parsedAst
			}(idl)
		}
	}

	wg.Wait()
	close(result)
	<-done

	if err != nil {
		return 0, err
	}
	if ast == nil {
		return 0, fmt.Errorf("no services found")
	}
	// TODO The parameters of the build request and the corresponding parameters are stored in the database
	var count int
	err = a.tm.Transaction(ctx, func(ctx context.Context) error {
		var module model.ModuleInfo
		var interfaceInfo model.InterfaceInfo
		for _, service := range ast.Services {
			module.ID = a.sid.GenString()
			module.BranchID = branchId
			module.ModuleName = service.Name
			module.InterfaceNum = int32(len(service.Functions))
			count = count + len(service.Functions)
			if err = a.db.CreateModule(ctx, &module); err != nil {
				return err
			}
			for _, function := range service.Functions {
				interfaceInfo.ID = a.sid.GenString()
				interfaceInfo.ModuleID = module.ID
				interfaceInfo.InterfaceName = function.Name
				method, url, err := interface_info.ParserMethodAndUrl(function)
				if err != nil {
					return err
				}
				interfaceInfo.Method = method
				interfaceInfo.URL = url
				parameter, err := interface_info.GetParameter(&structure, function)
				if err != nil {
					return err
				}
				interfaceInfo.Parameter = string(*parameter)
				response, ok := structure.Load(function.GetFunctionType().GetName())
				if !ok {
					interfaceInfo.Response = function.GetFunctionType().GetName()
				} else {
					responseBytes, err := sonic.Marshal(response)
					if err != nil {
						return err
					}
					interfaceInfo.Response = string(responseBytes)
				}
				if err = a.db.CreateInterface(ctx, &interfaceInfo); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewAnalyzeService(
	service *Service,
	queryRepo repository.QueryRepository,
	captchaRepo repository.CaptchaRepository,
) AnalyzeService {
	return &analyzeService{
		Service: service,
		db:      queryRepo,
		rdb:     captchaRepo,
	}
}
