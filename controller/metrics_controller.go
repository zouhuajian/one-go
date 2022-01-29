package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"one-go/internal"
)

// MetricsHierarchy ...
type MetricsHierarchy struct {
	appname string `json:"appname"`
	name    string `json:"name"`
}

// GetMetricsHierarchy
// @Router /api/v1/metrics/hierarchy
func GetMetricsHierarchy(c *gin.Context) {
	fmt.Printf("c=%v", c)
	internal.APIResponse(c, nil, MetricsHierarchy{
		name: "test",
	})
}
