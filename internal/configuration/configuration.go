package configuration

import (
	"desafio-backend/pkg/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var initialValues = map[string]string{
	"DB_USER":     "postgres",
	"DB_PASSWORD": "desafiobackend",
	"DB_HOST":     "postgres",
	"DB_NAME":     "postgres",
	"DB_SSL":      "disable",
	"DB_PORT":     "5432",
}

func GetDBConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		GetConfiguration("DB_HOST"),
		GetConfiguration("DB_PORT"),
		GetConfiguration("DB_USER"),
		GetConfiguration("DB_NAME"),
		GetConfiguration("DB_PASSWORD"),
		GetConfiguration("DB_SSL"),
	)
	logger.Info(fmt.Sprintf("Connection string"), dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}

func GetConfiguration(configuration string) string {
	return initialValues[configuration]
}
