package postgres

import (
	"fmt"
	"log/slog"

	"github.com/charmingruby/g3/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
	slog.Info("Connecting to database...")

	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.DatabaseName,
		cfg.DatabaseConfig.SSL,
	)
	dbDriver := "postgres"

	db, err := sqlx.Connect(dbDriver, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	slog.Info("Connected successfully to the database!")

	return db, nil
}
