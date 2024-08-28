package repository

import (
	"database/sql"
	"fmt"
	"github/desafio/models"
	"github/desafio/service"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	user     = "postgres"
	password = "desafio"
	dbname   = "postgres"
)

type Repo struct {
	db *sql.DB
}

var _ service.Repository = (*Repo)(nil)

func NewRepository() (service.Repository, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	return &Repo{db: db}, nil
}

func DBConnection() (*sql.DB, error) {
	host := os.Getenv("host")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
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

func (r *Repo) InsertGyroscopeData(data models.Gyroscope) error {

	insertQuery := `INSERT INTO public.gyroscope_data
	(macaddress, x_coord, y_coord, z_coord, data_timestamp, created)
	VALUES($1, $2, $3, $4, $5, $6);`

	sqlStatement, err := r.db.Prepare(insertQuery)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(data.MacAddress, data.X, data.Y, data.Z, data.Timestamp, time.Now().Unix())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}


func (r *Repo) InsertGPSData(data models.GPS) error {

	insertQuery := `INSERT INTO public.gps_data
	(macaddress, latitude, longitude, data_timestamp, created)
	VALUES($1, $2, $3, $4, $5);`

	sqlStatement, err := r.db.Prepare(insertQuery)
	if err != nil {
		println("gps insert error: ")
		log.Println(err.Error())
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(data.MacAddress, data.Latitude, data.Longitude, data.Timestamp, time.Now().Unix())
	if err != nil {
		println("gps insert error: ")
		log.Println(err.Error())
		return err
	}

	return nil
}


func (r *Repo) InsertPhoto(data models.Photo) error {

	insertQuery := `INSERT INTO public.photo
	(macaddress, photo, data_timestamp, created)
	VALUES($1, decode($2, 'base64'), $3, $4);`

	sqlStatement, err := r.db.Prepare(insertQuery)
	if err != nil {
		println("photo insert error: ")
		log.Println(err.Error())
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(data.MacAddress, data.Photo, data.Timestamp, time.Now().Unix())
	if err != nil {
		println("photo insert error: ")
		log.Println(err.Error())
		return err
	}

	return nil
}