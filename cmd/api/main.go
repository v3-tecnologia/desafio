package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/g3/config"
	"github.com/charmingruby/g3/internal/common/api/rest"
	"github.com/charmingruby/g3/internal/telemetry/database/postgres_repository"
	"github.com/charmingruby/g3/internal/telemetry/domain/usecase"
	v1 "github.com/charmingruby/g3/internal/telemetry/transport/rest/endpoint/v1"
	"github.com/charmingruby/g3/pkg/postgres"
	"github.com/charmingruby/g3/test/inmemory_repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	db, err := postgres.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE: %s", err.Error()))
		os.Exit(1)
	}

	router := gin.Default()
	rest.SetupCORS(router)

	initDependencies(router, db)

	server := rest.NewServer(router, "3000")

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	slog.Info("HTTP Server interruption received!")

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Server.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("GRACEFUL SHUTDOWN REST SERVER: %s", err.Error()))
		os.Exit(1)
	}

	slog.Info("Gracefully shutdown!")
}

func initDependencies(router *gin.Engine, db *sqlx.DB) {
	gpsRepo, err := postgres_repository.NewGPSPostgresRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	photoRepo := inmemory_repository.NewPhotoInMemoryRepository()

	gyroscopeRepo, err := postgres_repository.NewGyroscopePostgresRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	telemetryService := usecase.NewTelemetryUseCaseRegistry(gpsRepo, gyroscopeRepo, photoRepo)

	v1.NewHandler(router, &telemetryService).Register()
}
