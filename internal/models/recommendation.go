package models

// RecommendationRequest mirrors the structure of the Python RecommendationRequest model.
type RecommendationRequest struct {
	TrailID      string `json:"trail_id" binding:"required"`
	UserPathDesc string `json:"user_path_desc" binding:"required"`
}

// RecommendationResponse mirrors the structure of the Python RecommendationResponse model.
type RecommendationResponse struct {
	SafetyScore    int    `json:"safety_score"`
	Recommendation string `json:"recommendation"`
	Reasoning      string `json:"reasoning"`
	DataSource     string `json:"data_source"`
}
