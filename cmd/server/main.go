package main

import (
	"github.com/WuSKai403/hiking-app-golang/configs"
	"github.com/WuSKai403/hiking-app-golang/internal/database"
	"github.com/WuSKai403/hiking-app-golang/internal/server"
	"github.com/WuSKai403/hiking-app-golang/internal/services"
)

func main() {
	// Load configuration
	configs.LoadConfig()

	// Initialize GenAI client
	services.InitGenAI()

	// Connect to the database
	database.Connect(configs.AppConfig.MongoURI)
	defer database.Disconnect()

	// Setup and start the server
	router := server.NewRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
