package handlers

import (
	"net/http"

	"github.com/WuSKai403/hiking-app-golang/internal/models"
	"github.com/WuSKai403/hiking-app-golang/internal/services"
	"github.com/gin-gonic/gin"
)

// GetRecommendation handles the AI recommendation request.
func GetRecommendation(c *gin.Context) {
	var request models.RecommendationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// --- Data Fetching ---
	// Fetch real CWA data
	weatherData, err := services.GetCWADataForAI(request.TrailID)
	if err != nil {
		// Log the error but proceed with a fallback message for the AI
		// This makes the service more resilient if CWA fails
		weatherData = "無法獲取即時天氣資訊。"
	}

	// TODO: Replace with actual review data fetching
	mockReviewData := "近期評論 (2025/09/06): 很好的一條路線。 (2025/04/16): 車位充足，步道平緩良好。"

	// --- AI Recommendation ---
	recommendation, err := services.GetAiRecommendation(request, weatherData, mockReviewData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI recommendation: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendation)
}
