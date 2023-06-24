package models

import (
	"database/sql"
	"time"
)

type Authentication struct {
	ID       int
	Token    string
	UserID   int
	ExpiryAt time.Time
	IsActive bool `default:"true"`
}

func CreateAuthentication(db *sql.DB, auth *Authentication) error {
	query := "INSERT INTO authentication (token, user_id, expiry_at) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(query, auth.Token, auth.UserID, auth.ExpiryAt).Scan(&auth.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAuthentication(db *sql.DB, auth *Authentication) error {
	query := "UPDATE authentication SET token = $1, user_id = $2, expiry_at = $3, WHERE id = $4"
	_, err := db.Exec(query, auth.Token, auth.UserID, auth.ExpiryAt, auth.ID)
	if err != nil {
		return err
	}
	return nil
}
