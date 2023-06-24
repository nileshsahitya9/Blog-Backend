package models

import (
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
