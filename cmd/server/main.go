package main

import (
	"log"
	"os"

	"github.com/WuSKai403/hiking-app-golang/configs"
	"github.com/WuSKai403/hiking-app-golang/internal/database"
	"github.com/WuSKai403/hiking-app-golang/internal/server"
	"github.com/WuSKai403/hiking-app-golang/internal/services"
)

func main() {
	// Load configuration
	configs.LoadConfig()

	// Log all environment variables for debugging
	log.Println("--- All Environment Variables ---")
	for _, env := range os.Environ() {
		// Avoid logging sensitive information in production if possible
		// For this debugging purpose, we will log them.
		// You might want to filter out sensitive keys in a real production environment.
		log.Println(env)
	}
	log.Println("---------------------------------")

	// Log the Gemini API Key to verify it's loaded
	if configs.AppConfig.GeminiAPIKey == "" {
		log.Println("Warning: GEMINI_API_KEY environment variable not set or empty.")
	} else {
		// To avoid leaking the key, just confirm its presence
		log.Println("GEMINI_API_KEY is present.")
	}
	if configs.AppConfig.MongoURI == "" {
		log.Println("Warning: MONGO_URI environment variable not set or empty.")
	} else {
		log.Println("MONGO_URI is present.")
	}

	// Initialize GenAI client
	services.InitGenAI()

	// Connect to the database
	database.Connect(configs.AppConfig.MongoURI)
	defer database.Disconnect()

	// Setup and start the server
	router := server.NewRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
