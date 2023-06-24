package models

import (
	"time"
)

type Blog struct {
	ID        int
	UserID    int
	Title     string
	SubTitle  string
	Body      string
	ViewCount int
	CreatedAT time.Time
	UpdatedAt time.Time
}
