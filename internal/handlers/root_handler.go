package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Root handles the request to the root URL.
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hiking2025 API is running."})
}
