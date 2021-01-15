package entity

import "time"

// ReportInsertable to create new reported cases data
type ReportInsertable struct {
	UserID    int
	ImagePath string
}

// ReportedCases data from database
type ReportedCases struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Status       int       `json:"status"`
	ImagePath    string    `json:"image_path"`
	ImageDeleted bool      `json:"image_deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ReportStatus ...
type ReportStatus struct {
	Status string `json:"status"`
}
