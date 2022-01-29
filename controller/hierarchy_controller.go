package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"one-go/internal"
)

// UserController ...
type MetricsHierarchy struct {
	appname string `json:"appname"`
	name    string `json:"name"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GetMetricsHierarchy
// @Router /api/v1/metrics/hierarchy
func GetMetricsHierarchy(c *gin.Context) {
	fmt.Printf("c=%v", c)
	internal.APIResponse(c, nil, MetricsHierarchy{
		name: "test",
	})
}
