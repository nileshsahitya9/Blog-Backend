package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
	"github.com/nileshsahitya9/Blogs-Backend/internal/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

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
