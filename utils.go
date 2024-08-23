package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type db_table interface {
	decode([]byte) bool
	persist(*sql.DB) bool
}

func makeHandler(ctor func() db_table, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := make([]byte, r.ContentLength)
		r.Body.Read(content)
		ptr := ctor()

		log.Print(r)
		log.Println(string(content[:]))

		if !ptr.decode(content) {
			http.Error(w, "Invalid values", http.StatusBadRequest)
			return
		}

		if !ptr.persist(db) {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
	}
}

func connectDatabase(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
