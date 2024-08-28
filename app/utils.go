package main

import (
	"database/sql"
	"io"
	"net/http"
)

type db_table interface {
	decode([]byte) bool
	persist(*sql.DB) error
}

func makeHandler(ctor func() db_table, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		ptr := ctor()

		if !ptr.decode(content) {
			http.Error(w, "Invalid values", http.StatusBadRequest)
			return
		}

		if err := ptr.persist(db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func connectDatabase(db_type string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(db_type, dsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

func validateDevice(vals map[string]interface{}) (string, bool) {
	deviceID, valid := vals["deviceID"].(string)

	if valid {
		valid = (deviceID != "")
	}

	return deviceID, valid
}

func validateTimestamp(vals map[string]interface{}) (uint64, bool) {
	time, valid := vals["timestamp"].(float64)

	if valid {
		valid = (time == float64(uint64(time)))
	}

	return uint64(time), valid
}
