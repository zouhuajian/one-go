package api

import (
	"github.com/gin-gonic/gin"
	"github.com/one-go/one-web/controller"
)

func RegisterRouterV1(g *gin.RouterGroup) {
	// 加载html文件，即template包下所有文件
	//router.LoadHTMLGlob("template/*")

	metricsApi := g.Group("/metrics")
	metricsApi.GET("/hierarchy", controller.GetMetricHierarchy)
}
