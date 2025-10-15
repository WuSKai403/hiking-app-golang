package models

import "time"

// TrailDocument represents the structure for a trail document in MongoDB.
type TrailDocument struct {
	ID            int       `bson:"_id" json:"id"`
	Title         string    `bson:"title" json:"title"`
	URL           string    `bson:"url" json:"url"`
	Location      string    `bson:"location" json:"location"`
	Content       string    `bson:"content" json:"content"`
	LastUpdated   time.Time `bson:"last_updated" json:"last_updated"`
	Comments      []Comment `bson:"comments" json:"comments"`
	SafetyResults *Safety   `bson:"safety_results,omitempty" json:"safety_results,omitempty"`
}

// Comment represents a user comment on a trail.
type Comment struct {
	Date    time.Time `bson:"date" json:"date"`
	Author  string    `bson:"author" json:"author"`
	Content string    `bson:"content" json:"content"`
}

// Safety represents the AI-generated safety analysis for a trail.
type Safety struct {
	Score       int       `bson:"score" json:"score"`
	Summary     string    `bson:"summary" json:"summary"`
	GeneratedAt time.Time `bson:"generated_at" json:"generated_at"`
}
