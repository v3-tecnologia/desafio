package handlers

import (
	"desafio/models"
	"desafio/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

//Aqui estão todas as funções de handler

// Essa estrutura serve para encapsular uma esrutura de serviço que é abstraida pela interface IService foi feito dessa forma para permitir testes unitarios sem integração
type RequestHandle struct {
	serv service.IService
}

// Contrutor para a estrutura de handle
func NewRequestHandle(service *service.Service) *RequestHandle {
	if service == nil {
		return &RequestHandle{}
	}

	return &RequestHandle{serv: service}
}

// Handle da rota HealthCheck
func (rh *RequestHandle) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("up and running...")
}

// Handle da rota Gyroscope
func (rh *RequestHandle) GyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	var gyroscopeRequest models.GyroscopeRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&gyroscopeRequest)
	if err != nil {
		println("gyroscope handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gyroscopeRequest); err != nil {
		println("gyroscope handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rh.serv.ProcessGyroscopeData(gyroscopeRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handle da rota Gps
func (rh *RequestHandle) GpsHandler(w http.ResponseWriter, r *http.Request) {
	var gpsRequest models.GpsRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&gpsRequest)
	if err != nil {
		println("gps handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gpsRequest); err != nil {
		println("gps handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rh.serv.ProcessGpsData(gpsRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handle da rota Photo
func (rh *RequestHandle) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	var photoRequest models.PhotoRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&photoRequest)
	if err != nil {
		println("photo handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(photoRequest); err != nil {
		println("photo handler error: ")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rh.serv.ProcessPhotoData(photoRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
