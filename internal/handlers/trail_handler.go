package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/WuSKai403/hiking-app-golang/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTrail handles the HTTP request to retrieve a trail by its ID.
func GetTrail(c *gin.Context) {
	// Get id from URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trail ID format"})
		return
	}

	// Call the service to get the trail data
	trail, err := services.GetTrailByID(id)
	if err != nil {
		// Check if the error is "document not found"
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Trail not found"})
			return
		}
		// For any other errors, return a generic server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trail data"})
		return
	}

	// Return the trail data with a 200 OK status
	c.JSON(http.StatusOK, trail)
}

// GetAllTrailsSummary handles the HTTP request to retrieve a summary of all trails.
func GetAllTrailsSummary(c *gin.Context) {
	// Call the service to get all trail summaries
	trails, err := services.GetAllTrailsSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trails summary"})
		return
	}

	// Return the trails data with a 200 OK status
	c.JSON(http.StatusOK, trails)
}
