package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckHealth ...
func CheckHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
