package repositories

import (
	"database/sql"
	"github.com/ThalesMonteir0/desafio/internal/ports"
)

type telemetryRepository struct {
	db *sql.DB
}

func NewTelemetryRepository(db *sql.DB) ports.TelemetryRepository {
	return &telemetryRepository{
		db: db,
	}
}
