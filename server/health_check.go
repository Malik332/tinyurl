package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary server health check
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 {string} string	"health"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": HealthStatus,
	})
}
