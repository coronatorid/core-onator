package entity

import "time"

// TrackRequest ...
type TrackRequest struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Location ...
type Location struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LocationInsertable ...
type LocationInsertable struct {
	UserID int
	Lat    float64
	Long   float64
}
