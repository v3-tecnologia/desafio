package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GyroscopeRequest struct {
	Mac           string  `json:"macAddr"`
	X             float64 `json:"x"`
	Y             float64 `json:"y"`
	Z             float64 `json:"z"`
	UnixtimeStamp int64   `json:"timeStamp"`
}

type GpsRequest struct {
	Mac           string  `json:"macAddr"`
	Lat           float64 `json:"latitude"`
	Lon           float64 `json:"longitude"`
	UnixtimeStamp int64   `json:"timeStamp"`
}

type PhotoRequest struct {
	Mac           string `json:"macAddr"`
	UnixtimeStamp int64  `json:"timeStamp"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("up and running...")
}

func GyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	var gyroscopeRequest GyroscopeRequest

	err := json.NewDecoder(r.Body).Decode(&gyroscopeRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(gyroscopeRequest)
	return
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gps")
	var gpsRequest GpsRequest

	err := json.NewDecoder(r.Body).Decode(&gpsRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(gpsRequest)
	return
}

func PhotoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("photo")
}
