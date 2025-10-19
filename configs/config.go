package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application.
type Config struct {
	MongoURI     string `mapstructure:"MONGO_URI"`
	GeminiAPIKey string `mapstructure:"GEMINI_API_KEY"`
	CwaApiKey    string `mapstructure:"CWA_API_KEY"`
}

var AppConfig *Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() {
	AppConfig = &Config{} // Initialize the AppConfig

	// Set default values. This is crucial for Viper to recognize the keys
	// and be able to override them with environment variables even if the config file is missing.
	viper.SetDefault("MONGO_URI", "")
	viper.SetDefault("GEMINI_API_KEY", "")
	viper.SetDefault("CWA_API_KEY", "")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // Read in environment variables that match

	// Attempt to read the .env file. It's okay if it doesn't exist.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Info: .env file not found, relying on environment variables.")
		} else {
			log.Printf("Warning: Error reading .env file: %v", err)
		}
	}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}
}
