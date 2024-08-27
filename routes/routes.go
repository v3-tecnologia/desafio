package routes

import (
	"github/desafio/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/telemetry/gyroscope", handlers.GyroscopeData).Methods("POST")
	router.HandleFunc("/telemetry/gps", handlers.GPSData).Methods("POST")
	router.HandleFunc("/telemetry/photo", handlers.PhotoData).Methods("POST")
}
