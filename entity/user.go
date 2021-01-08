package entity

import "time"

// UserInsertable to create new user data
type UserInsertable struct {
	PhoneNumber string
}

// User data
type User struct {
	ID        int       `json:"id"`
	Phone     string    `json:"phone"`
	State     int       `json:"state"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
