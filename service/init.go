package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @Tags test-index
// @Success 200 {string} Hollowed
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
