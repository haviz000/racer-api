package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/haviz000/racer-api/models"
	"github.com/haviz000/racer-api/services"
	"github.com/haviz000/racer-api/utils"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	log.Println("methodnya", r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("decode error:", err)
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	log.Println("user name controller:", req.Username)
	log.Println("password controller:", req.Password)

	if req.Username == "" || req.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	if err := services.Login(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(req.Username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.LoginResponse{Token: token})
}
