package main

import (
	"log"
	"net/http"

	"github.com/haviz000/racer-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}
	routes.RegisterRoutes()

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
