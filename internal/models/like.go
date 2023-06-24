package models

import (
	"time"
)

type Like struct {
	ID        int
	BlogID    int
	UserID    int
	CreatedAT time.Time
	UpdatedAt time.Time
}
