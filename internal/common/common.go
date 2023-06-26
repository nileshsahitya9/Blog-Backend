package common

import "time"

type User struct {
	ID       int
	Email    string `json:"email"`
	Bio      string `json:"bio,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"userName"`
}

type OTP struct {
	ID     int
	OTP    string
	UserID int
}

type Authentication struct {
	ID       int
	Token    string
	UserID   int
	ExpiryAt time.Time
	IsActive bool `default:"true"`
}
