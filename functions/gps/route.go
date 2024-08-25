package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevenmiano/v3/internal/infra"
	database "github.com/kevenmiano/v3/internal/infra"
	repo "github.com/kevenmiano/v3/internal/repository/gps"
	route "github.com/kevenmiano/v3/internal/route/gps"
	service "github.com/kevenmiano/v3/internal/service/gps"
	"github.com/kevenmiano/v3/internal/shared"
	usecase "github.com/kevenmiano/v3/internal/usecase/gps"
)

func main() {

	logger := infra.NewLogger()

	gpsDatabase := database.NewGpsDatabase()

	gpsRepository := repo.NewGPSRepository(gpsDatabase)

	createGpsUseCase := usecase.NewCreateGpsUseCase(gpsRepository)

	gpsService := service.NewGpsService(logger, createGpsUseCase)

	handler := route.NewRoute(gpsService)

	lambdaHandler := shared.NewLambdaHandler(handler)

	lambda.Start(lambdaHandler)
}
