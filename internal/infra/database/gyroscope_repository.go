package database

import (
	"database/sql"
	domain "github.com/HaroldoFV/desafio/internal/domain/entity"
	_ "github.com/lib/pq"
)

type GyroscopeRepository struct {
	Db *sql.DB
}

func NewGyroscopeRepository(db *sql.DB) *GyroscopeRepository {
	return &GyroscopeRepository{Db: db}
}

func (r *GyroscopeRepository) Create(gyroscope *domain.Gyroscope) error {
	stmt, err := r.Db.Prepare("INSERT INTO gyroscopes (id, name, model, mac_address, x, y, z, timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gyroscope.GetID(), gyroscope.GetName(), gyroscope.GetModel(), gyroscope.GetMACAddress(), gyroscope.GetX(), gyroscope.GetY(), gyroscope.GetZ(), gyroscope.GetTimestamp())
	if err != nil {
		return err
	}
	return nil
}
