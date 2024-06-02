package handler

import (
	"context"
	"github.com/ByteBam/thirftbam/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type AnalyzeHandler struct {
	*Handler
	AnalyzeHandler service.AnalyzeService
}

func NewAnalyzeHandler(hander *Handler, analyzeHandler service.AnalyzeService) *AnalyzeHandler {
	return &AnalyzeHandler{
		Handler:        hander,
		AnalyzeHandler: analyzeHandler,
	}
}

func (h *AnalyzeHandler) Analyze(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, "analyze")
}
