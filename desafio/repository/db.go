package repository

import (
	"database/sql"
	"desafio/models"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func newPostgresConn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
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

func InsertGyroscopeData(data models.GyroscopeRequest) error {
	db, err := newPostgresConn()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO public.gyroscope
	(mac, xcoord, ycoord, zcoord, datatimestamp, created)
	VALUES($1, $2, $3, $4, $5, $6);`

	stmt, err := db.Prepare(query)
	if err != nil {
		println("gyroscope insert error: ")
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.X, data.Y, data.Z, data.UnixtimeStamp, time.Now().Unix())
	if err != nil {
		println("gyroscope insert error: ")
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
	VALUES($1, $2, $3, $4, $5);`

	stmt, err := db.Prepare(query)
	if err != nil {
		println("gps insert error: ")
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.Lat, data.Lon, data.UnixtimeStamp, time.Now().Unix())
	if err != nil {
		println("gps insert error: ")
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
	VALUES($1, decode($2, 'base64'), $3, $4);`

	stmt, err := db.Prepare(query)
	if err != nil {
		println("photo insert error: ")
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Mac, data.ImageBase64, data.UnixtimeStamp, time.Now().Unix())
	if err != nil {
		println("photo insert error: ")
		log.Println(err.Error())
		return err
	}

	return nil
}
