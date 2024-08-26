package repository

import (
	"database/sql"
	"desafio/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres-db"
)

func newPostgresConn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return db, nil
}

func InsertGyroscopeData(data models.GyroscopeRequest) error {
	db, err := newPostgresConn()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO public.gyroscope
	(mac, xcoord, ycoord, zcoord, datatimestamp, created)
	VALUES($1, $2, $3, $4, $5, now());`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.X, data.Y, data.Z, data.UnixtimeStamp)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func InsertGpsData(data models.GpsRequest) error {
	db, err := newPostgresConn()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO public.gps
	(mac, latitude, longitude, datatimestamp, created)
	VALUES($1, $2, $3, $4, now());`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.Lat, data.Lon, data.UnixtimeStamp)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func InsertPhotoData(data models.PhotoRequest) error {
	db, err := newPostgresConn()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO public.photo
	(mac, photo, datatimestamp, created)
	VALUES($1, decode($2, 'base64'), $3, now());`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.ImageBase64, data.UnixtimeStamp)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
