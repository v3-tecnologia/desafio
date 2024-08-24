package database

import (
	"database/sql"
	domain "github.com/HaroldoFV/desafio/internal/domain/entity"
	_ "github.com/lib/pq"
)

type GPSRepository struct {
	Db *sql.DB
}

func NewGPSRepository(db *sql.DB) *GPSRepository {
	return &GPSRepository{Db: db}
}

func (r *GPSRepository) Create(gps *domain.GPS) error {
	stmt, err := r.Db.Prepare("INSERT INTO gps (id, latitude, longitude, mac_address, timestamp) values ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gps.GetID(), gps.GetLatitude(), gps.GetLongitude(), gps.GetMACAddress(), gps.GetTimestamp())
	if err != nil {
		return err
	}
	return nil
}
