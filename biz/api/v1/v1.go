package v1

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *app.RequestContext, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: errorCodeMap[Success], Message: Success.Error(), Data: data}
	if _, ok := errorCodeMap[Success]; !ok {
		resp = Response{Code: 0, Message: "", Data: data}
	}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *app.RequestContext, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errorCodeMap[Success]; !ok {
		resp = Response{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}

func (e Error) Error() string {
	return e.Message
}
