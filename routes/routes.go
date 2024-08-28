package routes

import (
	"github/desafio/handlers"
	"github/desafio/repository"
	"github/desafio/service"
	"log"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	repository, err := repository.NewRepository()
	if err != nil{
		log.Fatal(err.Error())
	}
	handle := handlers.NewHandle(service.NewService(repository))
	router.HandleFunc("/telemetry/gyroscope", handle.GyroscopeData).Methods("POST")
	router.HandleFunc("/telemetry/gps", handle.GPSData).Methods("POST")
	router.HandleFunc("/telemetry/photo", handle.PhotoData).Methods("POST")
}
