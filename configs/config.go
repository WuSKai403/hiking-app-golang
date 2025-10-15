package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application.
type Config struct {
	MongoURI string `mapstructure:"MONGO_URI"`
}

var AppConfig *Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() {
	viper.AddConfigPath(".")    // Look for config in the current directory
	viper.SetConfigName(".env") // Name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name

	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: No .env file found, relying on environment variables: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}
}
