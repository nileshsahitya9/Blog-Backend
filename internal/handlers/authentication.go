package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
	"github.com/nileshsahitya9/Blogs-Backend/internal/services"
)

var jwtSecretToken string = os.Getenv("JWT_SECRET_KEY")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email string `json:"email"`
		OTP   int    `json:"otp"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := loginData.Email
	otp := loginData.OTP

	userId, userErr := services.GetUserByEmail(email)

	if userErr != nil {
		http.Error(w, userErr.Error(), http.StatusBadRequest)
		return
	}

	if !services.VerifyOTP(userId, otp) {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 3).Unix(), // Set token expiry to 3 days
	})

	secretKey := []byte(jwtSecretToken)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = services.StoreTokenInDatabase(userId, tokenString, time.Now().Add(time.Hour*24*3))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Login successful",
		Token:   tokenString,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user common.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.SaveUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	response := struct {
		Message string      `json:"message"`
		User    common.User `json:"user"`
	}{
		Message: "User registered successfully",
		User:    user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
