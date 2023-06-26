package services

import (
	"log"
	"time"

	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
	"github.com/nileshsahitya9/Blogs-Backend/internal/models"
)

func StoreTokenInDatabase(userID int, tokenString string, expiryTime time.Time) error {
	auth := &common.Authentication{
		UserID:   userID,
		Token:    tokenString,
		ExpiryAt: expiryTime,
	}
	err := models.CreateAuthentication(auth)

	if err != nil {
		log.Fatal("Failed to create a token in db:", err)
		return err
	}

	return nil

}
