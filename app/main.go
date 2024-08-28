package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Application expects 2 arguments <database config> <path to images directory>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	i := 0
	params := make([]string, 5)
	s := bufio.NewScanner(f)
	for s.Scan() {
		params[i] = s.Text()
		i++
	}

	cfg := mysql.Config{
		User:   params[0],
		Passwd: params[1],
		Net:    params[2],
		Addr:   params[3],
		DBName: params[4],
	}

	db, err := connectDatabase("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to the database!")
	}

	http.HandleFunc("POST /telemetry/gyroscope/", makeHandler(func() db_table { return &gyroscope{} }, db))
	http.HandleFunc("POST /telemetry/gps/", makeHandler(func() db_table { return &gps{} }, db))
	http.HandleFunc("POST /telemetry/photo/", makeHandler(func() db_table { return &photo{directory: os.Args[2] + "/"} }, db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
