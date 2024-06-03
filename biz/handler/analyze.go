package handler

import (
	"context"
	v1 "github.com/ByteBam/thirftbam/biz/api/v1"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"net/http"
)

type AnalyzeHandler struct {
	*Handler
	AnalyzeService service.AnalyzeService
}

func NewAnalyzeHandler(hander *Handler, analyzeService service.AnalyzeService) *AnalyzeHandler {
	return &AnalyzeHandler{
		Handler:        hander,
		AnalyzeService: analyzeService,
	}
}

// Analyze godoc
//
//	@Summary		analyze handler
//	@Description	analyze IDL file
//	@Tags			analyze
//	@Accept			json
//	@Produce		json
//	@Param			request	body		v1.AnalyzeRequest	true	"params"
//	@Success		200		{string}	string				v1.Response
//	@Router			/api/v1/analyze [post]
func (h *AnalyzeHandler) Analyze(ctx context.Context, c *app.RequestContext) {
	var req v1.AnalyzeRequest
	if err := c.BindAndValidate(&req); err != nil {
		logger.Error(err)
		v1.HandleError(c, http.StatusInternalServerError, v1.ErrParamError, err.Error())
		return
	}
	err := h.AnalyzeService.Download(ctx, &req)
	if err != nil {
		h.logger.WithContext(ctx).Error("AnalyzeService.Download error", zap.Error(err))
		v1.HandleError(c, http.StatusInternalServerError, v1.ErrDownloadError, err.Error())
		return
	}

	nums, err := h.AnalyzeService.Analyze(ctx, req.Owner)
	if err != nil {
		h.logger.WithContext(ctx).Error("AnalyzeService.Analyze error", zap.Error(err))
		v1.HandleError(c, http.StatusInternalServerError, v1.ErrServiceError, err.Error())
		return
	}

	go func(path string) {
		err = h.AnalyzeService.Delete(ctx, path)
		if err != nil {
			h.logger.WithContext(ctx).Error("AnalyzeService.DeleteFile error", zap.Error(err))
		}
	}(req.Owner)

	v1.HandleSuccess(c, nums)
}
