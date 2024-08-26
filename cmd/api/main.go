package main

import (
	"database/sql"
	"github.com/ThalesMonteir0/desafio/internal/handlers"
	"github.com/ThalesMonteir0/desafio/internal/repositories"
	"github.com/ThalesMonteir0/desafio/internal/routers"
	"github.com/ThalesMonteir0/desafio/internal/service"
	"github.com/ThalesMonteir0/desafio/pkg/database/postgresql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env.development file")
	}

	server := http.NewServeMux()

	db, err := postgresql.OpenConnDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	telemetryHandler := initDependencies(db)

	routers.TelemetryRoutes(server, telemetryHandler)

	if err = http.ListenAndServe(":5000", server); err != nil {
		log.Fatal(err.Error())
	}
}

func initDependencies(db *sql.DB) handlers.TelemetryHandler {
	telemetryRepository := repositories.NewTelemetryRepository(db)
	telemetryService := service.NewTelemetryService(telemetryRepository)
	return handlers.NewTelemetryHandler(telemetryService)
}
