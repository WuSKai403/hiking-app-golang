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
	// TODO: Replace these with actual data fetching logic similar to Python's
	// get_cwa_data_for_ai and get_hiking_reviews.
	mockWeatherData := "預報顯示：下午 1 點後有陣雨，夜間氣溫驟降至 10 度，風速 5 級。"
	mockReviewData := "近期評論 (2025/09/06): 很好的一條路線。 (2025/04/16): 車位充足，步道平緩良好。"

	// --- AI Recommendation ---
	recommendation, err := services.GetAiRecommendation(request, mockWeatherData, mockReviewData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI recommendation: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendation)
}
