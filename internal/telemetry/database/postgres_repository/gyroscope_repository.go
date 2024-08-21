package postgres_repository

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/database/postgres_repository/mapper"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	createGyroscope = "create gyroscope"
)

func gyroscopeQueries() map[string]string {
	return map[string]string{
		createGyroscope: `INSERT INTO gyroscopes
		(id, x_position, y_position, z_position)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
	}
}

func NewGyroscopePostgresRepository(db *sqlx.DB) (*GyroscopePostgresRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range gyroscopeQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "gyroscope", err)
		}

		stmts[queryName] = stmt
	}

	return &GyroscopePostgresRepository{
		db:    db,
		stmts: stmts,
	}, nil

}

type GyroscopePostgresRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *GyroscopePostgresRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "gyroscope")
	}

	return stmt, nil
}

func (r *GyroscopePostgresRepository) Store(e entity.Gyroscope) error {
	stmt, err := r.statement(createGyroscope)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainGyroscopeToPostgres(e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.XPosition,
		mappedEntity.YPosition,
		mappedEntity.ZPosition,
	); err != nil {
		return err
	}

	return nil
}
