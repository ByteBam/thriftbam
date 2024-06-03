package service

import (
	"context"
	"flag"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/util/config"
	"github.com/ByteBam/thirftbam/util/log"
	"github.com/ByteBam/thirftbam/util/sid"
	"testing"
)

func Test_analyzeService_Download(t *testing.T) {
	var envConf = flag.String("conf", "../config/config.yaml", "config file path")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	type fields struct {
		Service *Service
	}
	type args struct {
		ctx     context.Context
		request *v1.AnalyzeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_analyzeService_Download_1",
			fields: fields{
				Service: NewService(logger, sid.NewSid()),
			},
			args: args{
				ctx: context.Background(),
				request: &v1.AnalyzeRequest{
					Id:          "test",
					AccessToken: "770841cbc465a4b41de6a1d43c6dc933",
					Owner:       "liulangalliance",
					Repo:        "bam-backend",
					Path:        "idl",
					Ref:         "master",
				},
			},
			wantErr: false,
		},
		{
			name: "test_analyzeService_Download_2",
			fields: fields{
				Service: NewService(logger, sid.NewSid()),
			},
			args: args{
				ctx: context.Background(),
				request: &v1.AnalyzeRequest{
					Id:          "test",
					AccessToken: "770841cbc465a4b41de6a1d43c6dc933",
					Owner:       "liulangallian",
					Repo:        "bam-backend",
					Path:        "idl",
					Ref:         "master",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyzeService{
				Service: tt.fields.Service,
			}
			if err := a.Download(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_analyzeService_Analyze(t *testing.T) {
	var envConf = flag.String("conf", "../config/config.yaml", "config file path")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	type fields struct {
		Service *Service
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_analyzeService_Analyze_1",
			fields: fields{
				Service: NewService(logger, sid.NewSid()),
			},
			args: args{
				ctx:  context.Background(),
				path: "liulangalliance",
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyzeService{
				Service: tt.fields.Service,
			}
			got, err := a.Analyze(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Analyze() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_analyzeService_Delete(t *testing.T) {
	var envConf = flag.String("conf", "../config/config.yaml", "config file path")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	type fields struct {
		Service *Service
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_analyzeService_Delete_1",
			fields: fields{
				Service: NewService(logger, sid.NewSid()),
			},
			args: args{
				ctx:  context.Background(),
				path: "liulangalliance",
			},
			wantErr: false,
		},
		{
			name: "test_analyzeService_Delete_2",
			fields: fields{
				Service: NewService(logger, sid.NewSid()),
			},
			args: args{
				ctx:  context.Background(),
				path: "liulangallian",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyzeService{
				Service: tt.fields.Service,
			}
			if err := a.Delete(tt.args.ctx, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
