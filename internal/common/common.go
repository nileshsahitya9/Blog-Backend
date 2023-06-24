package common

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
