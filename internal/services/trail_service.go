package services

import (
	"context"
	"log"
	"time"

	"github.com/WuSKai403/hiking-app-golang/internal/database"
	"github.com/WuSKai403/hiking-app-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTrailByID retrieves a single trail document from the database by its ID.
func GetTrailByID(id int) (*models.TrailDocument, error) {
	var trail models.TrailDocument
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// The filter to find the document by its _id field
	filter := bson.M{"_id": id}

	// Find the document
	err := database.Trails.FindOne(ctx, filter).Decode(&trail)
	if err != nil {
		// This error could be mongo.ErrNoDocuments, which should be handled by the caller
		return nil, err
	}

	return &trail, nil
}

// GetAllTrailsSummary retrieves a list of trail summaries.
func GetAllTrailsSummary() ([]models.TrailSummary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"is_valid": true}
	projection := bson.M{
		"_id":        1,
		"name":       1,
		"location":   1,
		"difficulty": 1,
		"reviews":    1, // We need reviews to calculate the count
	}

	cursor, err := database.Trails.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var summaries []models.TrailSummary
	for cursor.Next(ctx) {
		var result struct {
			ID         int             `bson:"_id"`
			Name       string          `bson:"name"`
			Location   string          `bson:"location"`
			Difficulty string          `bson:"difficulty"`
			Reviews    []models.Review `bson:"reviews"`
		}
		if err := cursor.Decode(&result); err != nil {
			// Log the error but continue processing other documents
			log.Printf("Error decoding trail summary: %v", err)
			continue
		}

		summary := models.TrailSummary{
			ID:          result.ID,
			Name:        result.Name,
			Location:    result.Location,
			Difficulty:  result.Difficulty,
			ReviewCount: len(result.Reviews),
		}
		summaries = append(summaries, summary)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return summaries, nil
}
