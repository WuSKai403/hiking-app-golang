package models

import "time"

// Review represents a user review, mirroring the ReviewModel in Python.
type Review struct {
	UserID   string    `bson:"user_id" json:"user_id"`
	Username string    `bson:"username" json:"username"`
	Date     time.Time `bson:"review_date" json:"review_date"`
	Content  string    `bson:"content" json:"content"`
}

// TrailDocument represents the structure for a trail document in MongoDB,
// mirroring the TrailDocument Pydantic model.
type TrailDocument struct {
	ID                 int       `bson:"_id" json:"id"`
	Name               string    `bson:"name,omitempty" json:"name,omitempty"`
	Description        string    `bson:"description,omitempty" json:"description,omitempty"`
	Location           string    `bson:"location,omitempty" json:"location,omitempty"`
	Difficulty         string    `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	TrailType          string    `bson:"trail_type,omitempty" json:"trail_type,omitempty"`
	Distance           float64   `bson:"distance,omitempty" json:"distance,omitempty"`
	Altitude           string    `bson:"altitude,omitempty" json:"altitude,omitempty"`
	AltitudeDifference int       `bson:"altitude_difference,omitempty" json:"altitude_difference,omitempty"`
	Duration           string    `bson:"duration,omitempty" json:"duration,omitempty"`
	Pavement           string    `bson:"pavement,omitempty" json:"pavement,omitempty"`
	GPXUrl             string    `bson:"gpx_url,omitempty" json:"gpx_url,omitempty"`
	LastScrapedAt      time.Time `bson:"last_scraped_at" json:"last_scraped_at"`
	Reviews            []Review  `bson:"reviews" json:"reviews"`
	IsValid            bool      `bson:"is_valid" json:"is_valid"`
}

// TrailSummary represents a subset of TrailDocument for list views.
type TrailSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Difficulty  string `json:"difficulty"`
	ReviewCount int    `json:"review_count"`
}
