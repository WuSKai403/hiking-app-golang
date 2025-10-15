package server

import (
	"time"

	"github.com/WuSKai403/hiking-app-golang/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a Gin router.
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://hiking2025-front.pages.dev", "http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Root and health check endpoints
	router.GET("/", handlers.Root)
	router.GET("/ping", handlers.Ping)

	// Group routes for v2
	apiV2 := router.Group("/api/v2")
	{
		// Trail routes
		apiV2.GET("/trails", handlers.GetAllTrailsSummary)
		apiV2.GET("/trails/:id", handlers.GetTrail)

		// CWA routes
		apiV2.GET("/cwa/:trail_id", handlers.GetCWAData)

		// Recommendation routes
		apiV2.POST("/recommendation", handlers.GetRecommendation)
	}

	return router
}
