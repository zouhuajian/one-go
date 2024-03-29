package router

import (
	"github.com/gin-gonic/gin"
	"github.com/one-go/one-web/controller"
	"github.com/one-go/one-web/router/api"
)

// HttpServer ...
type HttpServer struct {
	GinEngine *gin.Engine
}

func InitRouter() *HttpServer {
	server := new(HttpServer)
	gin.SetMode(gin.DebugMode)

	server.GinEngine = gin.Default()
	registerBaseAPI(server)

	apiGroupV1 := server.GinEngine.Group("/api/v1")
	api.RegisterRouterV1(apiGroupV1)
	return server
}

// registerBaseAPI ...
// /health
func registerBaseAPI(server *HttpServer) {
	server.GinEngine.GET("/", controller.CheckHealth)
}
