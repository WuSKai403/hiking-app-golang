package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Trails *mongo.Collection
)

// Connect initializes the connection to MongoDB.
func Connect(uri string) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the primary
	if err := Client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")

	// Initialize collections
	Trails = Client.Database("hiking_db").Collection("trails")
}

// Disconnect closes the MongoDB connection.
func Disconnect() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		log.Println("Connection to MongoDB closed.")
	}
}
