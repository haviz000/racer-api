package routes

import (
	"net/http"

	"github.com/haviz000/racer-api/controllers"
	"github.com/haviz000/racer-api/middlewares"
)

func RegisterRoutes() {
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/race-test", middlewares.JWTAuth(controllers.RaceTestController))
}
