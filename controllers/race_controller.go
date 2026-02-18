package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/haviz000/racer-api/models"
	"github.com/haviz000/racer-api/services"
)

func RaceTestController(w http.ResponseWriter, r *http.Request) {
	var req models.RaceRequest
	json.NewDecoder(r.Body).Decode(&req)

	response := services.ExecuteRaceTest(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
