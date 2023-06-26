package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
	helper "github.com/nileshsahitya9/Blogs-Backend/internal/helper/mail"
	"github.com/nileshsahitya9/Blogs-Backend/internal/models"
)

func SaveUser(user *common.User) error {

	exists, err := models.CheckEmailExists(user.Email)
	if err != nil {
		log.Fatal("Failed to check email existence:", err)
		return err
	}
	if exists {
		return errors.New("Email already exists")
	}

	var otp common.OTP

	user.Username = generateUsername()

	// Generate an OTP
	otp.OTP = generateOTP()

	userErr := models.CreateUser(user)

	if userErr != nil {
		log.Fatal("Failed to create user:", userErr)
		return errors.New("Failed to create user:")
	}

	otp.UserID = user.ID

	otpErr := models.CreateOTP(&otp)

	if otpErr != nil {
		log.Fatal("Failed to create OTP:", otpErr)
		return errors.New("Failed to create OTP")
	}

	mailErr := helper.SendOTP(user.Email, otp.OTP)

	if mailErr != nil {
		log.Fatal("Failed to send OTP on mail", err)
		return errors.New("Failed to send OTP on mail")
	}

	return nil
}

func generateUsername() string {
	rand.Seed(time.Now().UnixNano())
	adjectives := []string{"happy", "brave", "clever", "kind", "smart", "friendly"}
	nouns := []string{"cat", "dog", "bird", "lion", "tiger", "elephant"}
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s_%s_%d", adjective, noun, timestamp)
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return fmt.Sprintf("%06d", otp)
}

func GetUserByEmail(email string) (int, error) {

	userId := models.GetUserIDByEmail(email)

	if userId < 0 {
		return -1, errors.New("User not exists")
	}

	return userId, nil

}
