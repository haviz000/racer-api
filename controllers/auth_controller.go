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

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	log.Println("user name controller", req.Username, "kocak", req.Password)
	if err := services.Login(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(req.Username)
	json.NewEncoder(w).Encode(models.LoginResponse{Token: token})

}
