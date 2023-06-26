package models

import (
	"github.com/nileshsahitya9/Blogs-Backend/db"
	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
)

func CreateAuthentication(auth *common.Authentication) error {
	query := "INSERT INTO authentications (token, user_id, expiry_at) VALUES ($1, $2, $3) RETURNING id"
	err := db.DB.QueryRow(query, auth.Token, auth.UserID, auth.ExpiryAt).Scan(&auth.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAuthentication(auth *common.Authentication) error {
	query := "UPDATE authentications SET token = $1, user_id = $2, expiry_at = $3, WHERE id = $4"
	_, err := db.DB.Exec(query, auth.Token, auth.UserID, auth.ExpiryAt, auth.ID)
	if err != nil {
		return err
	}
	return nil
}
