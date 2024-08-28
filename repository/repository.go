package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	user     = "postgres"
	password = "desafio"
	dbname   = "postgres"
)

func DBConnection() (*sql.DB, error) {
	host := os.Getenv("host")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println("postgres connection error: ")
		log.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		println("postgres connection error: ")
		log.Println(err.Error())
		return nil, err
	}

	return db, nil
}
