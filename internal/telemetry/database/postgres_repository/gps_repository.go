package postgres_repository

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/database/postgres_repository/mapper"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	createGPS = "create gps"
)

func gpsQueries() map[string]string {
	return map[string]string{
		createGPS: `INSERT INTO gps
		(id, latitude, longitude)
		VALUES ($1, $2, $3)
		RETURNING *`,
	}
}

func NewGPSPostgresRepository(db *sqlx.DB) (*GPSPostgresRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range gpsQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "gps", err)
		}

		stmts[queryName] = stmt
	}

	return &GPSPostgresRepository{
		db:    db,
		stmts: stmts,
	}, nil

}

type GPSPostgresRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *GPSPostgresRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "gps")
	}

	return stmt, nil
}

func (r *GPSPostgresRepository) Store(e entity.GPS) error {
	stmt, err := r.statement(createGPS)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainGPSToPostgres(e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.Latitude,
		mappedEntity.Longitude,
	); err != nil {
		return err
	}

	return nil
}
