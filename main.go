package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

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

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "v3",
	}

	db, err := connectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to the database!")
	}

	http.HandleFunc("POST /telemetry/gyroscope/", makeHandler(func() db_table { return &gyroscope{} }, db))
	//http.HandleFunc("POST /telemetry/gps/", makeHandler(func() db_table { return &gps{} }, db))
	//http.HandleFunc("POST /telemetry/photo/", makeHandler(func() db_table { return &photo{} }, db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
