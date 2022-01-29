package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one-go/internal/code"
)

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// APIResponse ....
func APIResponse(Ctx *gin.Context, err error, data interface{}) {
	code, message := code.DecodeErr(err)
	Ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
}
