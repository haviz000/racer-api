package routes

import (
	"net/http"

	"github.com/haviz000/racer-api/controllers"
	"github.com/haviz000/racer-api/middlewares"
)

func RegisterRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controllers.LoginController)
	mux.HandleFunc("/race-test", middlewares.JWTAuth(controllers.RaceTestController))

	handler := middlewares.CORS(mux)
	http.Handle("/", handler)
}
