//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	"github.com/HaroldoFV/desafio/internal/domain"
	"github.com/HaroldoFV/desafio/internal/gateway"
	"github.com/HaroldoFV/desafio/internal/infra/aws"
	"github.com/HaroldoFV/desafio/internal/infra/database"
	"github.com/HaroldoFV/desafio/internal/infra/web"
	"github.com/HaroldoFV/desafio/internal/infra/web/webserver"
	"github.com/HaroldoFV/desafio/internal/usecase"
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
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
		provideAWSConfig,
		NewRekognitionClient,
		provideRekognitionFaceRecognizer,
		provideCreateGyroscopeUseCase,
		provideCreateGPSUseCase,
		provideCreatePhotoUseCase,
		provideGyroscopeHandler,
		provideGPSHandler,
		providePhotoHandler,
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

func provideAWSConfig(conf *configs.Conf) (awssdk.Config, error) {
	return awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(conf.AWSRegion),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			conf.AWSAccessKeyID,
			conf.AWSSecretAccessKey,
			"",
		)),
	)
}

func NewRekognitionClient(cfg awssdk.Config) *rekognition.Client {
	return rekognition.NewFromConfig(cfg)
}

func provideRekognitionFaceRecognizer(client *rekognition.Client, config *configs.Conf) gateway.FaceRecognizerInterface {
	return aws.NewRekognitionFaceRecognizer(client, config.AWSRekognitionCollectionID)
}

func provideCreateGyroscopeUseCase(repo domain.GyroscopeRepositoryInterface) usecase.CreateGyroscopeUseCaseInterface {
	return usecase.NewCreateGyroscopeUseCase(repo)
}

func provideCreateGPSUseCase(repo domain.GPSRepositoryInterface) usecase.CreateGPSUseCaseInterface {
	return usecase.NewCreateGPSUseCase(repo)
}

func provideCreatePhotoUseCase(repo domain.PhotoRepositoryInterface, faceRecognizer gateway.FaceRecognizerInterface, config *configs.Conf) usecase.CreatePhotoUseCaseInterface {
	return usecase.NewCreatePhotoUseCase(repo, faceRecognizer, config)
}

func provideGyroscopeHandler(useCase usecase.CreateGyroscopeUseCaseInterface) *web.GyroscopeHandler {
	return web.NewGyroscopeHandler(useCase)
}

func provideGPSHandler(useCase usecase.CreateGPSUseCaseInterface) *web.GPSHandler {
	return web.NewGPSHandler(useCase)
}

func providePhotoHandler(useCase usecase.CreatePhotoUseCaseInterface) *web.PhotoHandler {
	return web.NewPhotoHandler(useCase)
}
