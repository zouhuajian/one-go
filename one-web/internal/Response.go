package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/one-go/one-web/internal/code"
	"net/http"
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
