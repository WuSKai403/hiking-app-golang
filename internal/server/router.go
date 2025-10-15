package server

import (
	"github.com/WuSKai403/hiking-app-golang/internal/handlers"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a Gin router.
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/ping", handlers.Ping)

	// Here we can group routes for v2
	// apiV2 := router.Group("/api/v2")
	// {
	// 	// Add trail routes here later
	// }

	return router
}
