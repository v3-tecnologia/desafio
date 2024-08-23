package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

type db_table interface {
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

func (g *gyroscope) decode(data []byte) bool {
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	deviceID, valid_device := m["deviceID"].(string)
	timestamp := m["timestamp"].(float64)
	valid_time := timestamp == float64(uint64(timestamp))
	x, valid_x := m["x"].(float64)
	y, valid_y := m["y"].(float64)
	z, valid_z := m["z"].(float64)

	g.deviceID = deviceID
	g.timestamp = uint64(timestamp)
	g.x = x
	g.y = y
	g.z = z

	return valid_device && valid_time && valid_x && valid_y && valid_z
}

func (g *gyroscope) persist() {
}

func makeHandler(ctor func() db_table) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := make([]byte, r.ContentLength)
		r.Body.Read(content)
		ptr := ctor()

		log.Print(r)
		log.Println(string(content[:]))

		if !ptr.decode(content) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ptr.persist()
	}
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "v3",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected to the database!")

	http.HandleFunc("POST /telemetry/gyroscope/", makeHandler(func() db_table { return &gyroscope{} }))
	//http.HandleFunc("POST /telemetry/gps/", makeHandler(func() db_table { return &gps{} }))
	//http.HandleFunc("POST /telemetry/photo/", makeHandler(func() db_table { return &photo }))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
