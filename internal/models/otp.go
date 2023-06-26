package models

import (
	"log"

	"github.com/nileshsahitya9/Blogs-Backend/db"
	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
)

func CreateOTP(otp *common.OTP) error {
	query := `
		INSERT INTO otps (otp, user_id)
		VALUES ($1, $2)
		RETURNING id
	`
	err := db.DB.QueryRow(query, otp.OTP, otp.UserID).Scan(&otp.ID)
	if err != nil {
		return err
	}
	return nil
}

func VerifyOTP(userID, otp int) (int, error) {
	query := "SELECT otp FROM otps WHERE user_id = $1 AND created_at >= NOW() - INTERVAL '15 minutes' ORDER BY created_at DESC LIMIT 1"
	row := db.DB.QueryRow(query, userID)

	var OTP int
	err := row.Scan(&OTP)
	if err != nil {
		log.Fatalln("Failed to fetch OTP from the database:", err)
		return -1, err
	}
	return OTP, nil
}
