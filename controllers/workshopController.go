package controllers

import (
	"encoding/json"
	"microservicelistworkshops/config"
	"microservicelistworkshops/models"
	"net/http"
)

// GetAllWorkshops
// @Summary Obtiene todos los Workshops
// @Description Devuelve una lista JSON con todos los Workshops existentes en la base de datos
// @Tags Workshops
// @Produce json
// @Success 200 {array} models.Workshop
// @Failure 500 {string} string "Internal Server Error"
// @Router /workshops [get]
func GetAllWorkshops(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDatabase()

	workshops, err := models.GetAllWorkshops(db)
	if err != nil {
		http.Error(w, "Error retrieving workshops", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workshops)
}

// HealthCheck
// @Summary Verifica el estado del microservicio
// @Description Retorna un mensaje que indica que el microservicio está en funcionamiento
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Microservice is up and running"})
}
