package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a handler function for the /ping endpoint.
// It returns a simple JSON response to indicate the server is alive.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
