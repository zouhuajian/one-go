package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/one-go/one-web/internal"
)

// CheckHealth ...
func CheckHealth(c *gin.Context) {
	internal.APIResponse(c, nil, "ok")
}
