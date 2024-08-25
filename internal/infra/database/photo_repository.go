package database

import (
	"database/sql"
	domain "github.com/HaroldoFV/desafio/internal/domain/entity"
	_ "github.com/lib/pq"
)

type PhotoRepository struct {
	Db *sql.DB
}

func NewPhotoRepository(db *sql.DB) *PhotoRepository {
	return &PhotoRepository{Db: db}
}

func (r *PhotoRepository) Create(photo *domain.Photo) error {
	stmt, err := r.Db.Prepare("INSERT INTO photos (id, file_path, mac_address, timestamp) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(photo.GetID(), photo.GetFilePath(), photo.GetMACAddress(), photo.GetTimestamp())
	if err != nil {
		return err
	}
	return nil
}
