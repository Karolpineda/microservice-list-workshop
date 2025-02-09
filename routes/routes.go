package routes

import (
	"microservicelistworkshops/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Ruta Health
	r.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	r.HandleFunc("/workshops", controllers.GetAllWorkshops).Methods("GET")

}
