package main

import (
	"log"
	"net/http"
)

type db_table interface {
	validate([]byte) bool
	decode([]byte) bool
	persist()
}

type gyroscope struct {
	x, y, z   float64
	timestamp uint64
	deviceID  string
}

//type gps struct {
//latitude, longitude float64
//timestamp           uint64
//deviceID            string
//}

//type photo struct {
//image     []byte
//timestamp uint64
//deviceID  string
//}

func (g *gyroscope) validate([]byte) bool {

	return true
}

func (g *gyroscope) decode([]byte) bool {
	return true
}

func (g *gyroscope) persist() {
}

func makeHandler(ctor func() db_table) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := make([]byte, r.ContentLength)
		r.Body.Read(content)
		ptr := ctor()

		if ptr.validate(content) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if ptr.decode(content) {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}

		ptr.persist()
	}
}

func main() {
	http.HandleFunc("POST /telemetry/gyroscope/", makeHandler(func() db_table { return &gyroscope{} }))
	//http.HandleFunc("POST /telemetry/gps/", makeHandler(func() db_table { return &gps{} }))
	//http.HandleFunc("POST /telemetry/photo/", makeHandler(func() db_table { return &photo }))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
