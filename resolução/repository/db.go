package repository

import (
	"database/sql"
	"desafio/models"
	"desafio/service"
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

// Estrutura que contém a conexão com banco de dados
type Repository struct {
	db *sql.DB
}

// Validação que a estrutura assina a interface
var _ service.IRepository = (*Repository)(nil)

// Contrutor da estrutura de Repository
func NewRepository() (service.IRepository, error) {
	db, err := newPostgresConn()
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// Cria conexão com o banco de dados
func newPostgresConn() (*sql.DB, error) {
	host := os.Getenv("host")
	if host == "" {
		host = "localhost"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

// Função que insere os dados de Gyroscope no banco
func (r *Repository) InsertGyroscopeData(data models.GyroscopeRequest) error {

	query := `INSERT INTO public.gyroscope
	(mac, xcoord, ycoord, zcoord, datatimestamp, created)
	VALUES($1, $2, $3, $4, $5, $6);`

	stmt, err := r.db.Prepare(query)
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

// Função que insere os dados de Gps no banco
func (r *Repository) InsertGpsData(data models.GpsRequest) error {

	query := `INSERT INTO public.gps
	(mac, latitude, longitude, datatimestamp, created)
	VALUES($1, $2, $3, $4, $5);`

	stmt, err := r.db.Prepare(query)
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

// Função que insere os dados de Photo no banco
func (r *Repository) InsertPhotoData(data models.PhotoRequest) error {

	query := `INSERT INTO public.photo
	(mac, photo, datatimestamp, created)
	VALUES($1, decode($2, 'base64'), $3, $4);`

	stmt, err := r.db.Prepare(query)
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
