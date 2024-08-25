//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	"github.com/HaroldoFV/desafio/internal/domain"
	"github.com/HaroldoFV/desafio/internal/infra/database"
	"github.com/HaroldoFV/desafio/internal/infra/web"
	"github.com/HaroldoFV/desafio/internal/infra/web/webserver"
	"github.com/HaroldoFV/desafio/internal/usecase"
	"github.com/google/wire"
)

type Application struct {
	WebServer        *webserver.WebServer
	GyroscopeHandler *web.GyroscopeHandler
	GPSHandler       *web.GPSHandler
	PhotoHandler     *web.PhotoHandler
}

func InitializeApplication(config *configs.Conf) (*Application, error) {
	wire.Build(
		provideDB,
		provideGyroscopeRepository,
		provideGPSRepository,
		providePhotoRepository,
		providePhotoStoragePath,
		usecase.NewCreateGyroscopeUseCase,
		usecase.NewCreateGPSUseCase,
		usecase.NewCreatePhotoUseCase,
		web.NewGyroscopeHandler,
		web.NewGPSHandler,
		web.NewPhotoHandler,
		provideWebServer,
		wire.Struct(new(Application), "*"),
	)
	return nil, nil
}

func provideDB(config *configs.Conf) (*sql.DB, error) {
	return sql.Open(config.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
}

func provideWebServer(config *configs.Conf) *webserver.WebServer {
	return webserver.NewWebServer(":" + config.WebServerPort)
}

func provideGyroscopeRepository(db *sql.DB) domain.GyroscopeRepositoryInterface {
	return database.NewGyroscopeRepository(db)
}

func provideGPSRepository(db *sql.DB) domain.GPSRepositoryInterface {
	return database.NewGPSRepository(db)
}

func providePhotoRepository(db *sql.DB) domain.PhotoRepositoryInterface {
	return database.NewPhotoRepository(db)
}

func providePhotoStoragePath(config *configs.Conf) string {
	return config.PhotoStoragePath
}
