package postgres_repository

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/database/postgres_repository/mapper"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	createPhoto = "create photo"
)

func photoQueries() map[string]string {
	return map[string]string{
		createPhoto: `INSERT INTO photos
		(id, image_url, is_recognized, amount_of_faces_detected, confidence_mean)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *`,
	}
}

func NewPhotoPostgresRepository(db *sqlx.DB) (*PhotoPostgresRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range photoQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "photos", err)
		}

		stmts[queryName] = stmt
	}

	return &PhotoPostgresRepository{
		db:    db,
		stmts: stmts,
	}, nil

}

type PhotoPostgresRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PhotoPostgresRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "photos")
	}

	return stmt, nil
}

func (r *PhotoPostgresRepository) Store(e entity.Photo) error {
	stmt, err := r.statement(createPhoto)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainPhotoToPostgres(e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.ImageURL,
		mappedEntity.IsRecognized,
		mappedEntity.AmountOfFacesDetected,
		mappedEntity.ConfidenceMean,
	); err != nil {
		return err
	}

	return nil
}
