package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Responses 响应结果
func Responses(ctx *gin.Context, code int, msg string, data interface{}) {
	resp := ResponseData{
		Code: code,
		Data: data,
	}
	if msg != "" {
		resp.Msg = msg
	} else {
		resp.Msg = ErrorMap[code]
	}

	ctx.JSON(http.StatusOK, resp)
}
