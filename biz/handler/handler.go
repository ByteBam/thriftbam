// Code generated by hertz generator.

package handler

import "github.com/ByteBam/thirftbam/pkg/util/log"

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}
