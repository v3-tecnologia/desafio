package handlers

import (
	"desafio/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("up and running...")
}

func GyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	var gyroscopeRequest models.GyroscopeRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&gyroscopeRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gyroscopeRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	var gpsRequest models.GpsRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&gpsRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gpsRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func PhotoHandler(w http.ResponseWriter, r *http.Request) {
	var photoRequest models.PhotoRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&photoRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(photoRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
