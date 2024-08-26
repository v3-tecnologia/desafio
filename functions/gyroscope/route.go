package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevenmiano/v3/internal/infra"
	repo "github.com/kevenmiano/v3/internal/repository/gyroscope"
	route "github.com/kevenmiano/v3/internal/route/gyroscope"
	service "github.com/kevenmiano/v3/internal/service/gyroscope"
	"github.com/kevenmiano/v3/internal/shared"
	usecase "github.com/kevenmiano/v3/internal/usecase/gyroscope"
)

func main() {

	database := infra.NewGyroscopeDatabase()

	logger := infra.NewLogger()

	gyroscopeRepository := repo.NewGyroscopeRepository(database)

	gyroscopeUseCase := usecase.NewCreateGyroscopeUseCase(gyroscopeRepository)

	gyroscopeService := service.NewGyroscopeService(logger, gyroscopeUseCase)

	handler := route.NewRoute(gyroscopeService)

	lambdaHandler := shared.NewLambdaHandler(handler)

	lambda.Start(lambdaHandler)
}
