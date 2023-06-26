package services

import (
	"github.com/nileshsahitya9/Blogs-Backend/internal/models"
)

func VerifyOTP(userID, OTP int) bool {

	storedOTP, err := models.VerifyOTP(userID, OTP)

	if err != nil {
		return false
	}

	if storedOTP != OTP {
		return false
	}
	return true

}
