package handlers

import (
	"net/http"

	"github.com/WuSKai403/hiking-app-golang/internal/services"
	"github.com/gin-gonic/gin"
)

// GetCWAData handles the request for CWA data for a specific trail.
func GetCWAData(c *gin.Context) {
	trailID := c.Param("trail_id")
	if trailID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Trail ID is required",
		})
		return
	}

	cwaSummary, err := services.GetCWADataForAI(trailID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, cwaSummary)
}
