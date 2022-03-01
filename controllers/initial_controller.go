package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ping handler functions for endpoints
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}

// welcome handler functions for endpoints
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to ContactBook API",
	})
}
