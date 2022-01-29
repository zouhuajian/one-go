package controller

import (
	"github.com/gin-gonic/gin"
	"one-go/internal"
)

// CheckHealth ...
func CheckHealth(c *gin.Context) {
	internal.APIResponse(c, nil, "ok")
}
