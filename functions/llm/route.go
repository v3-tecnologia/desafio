package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevenmiano/v3/internal/infra"
	rekognitionRepository "github.com/kevenmiano/v3/internal/repository/rekognition"
	route "github.com/kevenmiano/v3/internal/route/llm"
	service "github.com/kevenmiano/v3/internal/service/rekognition"
	"github.com/kevenmiano/v3/internal/shared"
	usecase "github.com/kevenmiano/v3/internal/usecase/rekognition"
)

func main() {

	logger := infra.NewLogger()

	rekognition := infra.NewRekognition()

	newRekognitionRepository := rekognitionRepository.NewRekognitionRepository(rekognition)

	createIndexFaceUseCase := usecase.NewCreateIndexFaceUseCase(newRekognitionRepository)

	findFaceImageUseCase := usecase.NewSearchFaceImageUseCase(newRekognitionRepository)

	rekognitionService := service.NewRekognitionService(logger, createIndexFaceUseCase, findFaceImageUseCase)

	handler := route.NewRoute(rekognitionService)

	eventHandler := shared.NewEventLambdaHandler(handler)

	lambda.Start(eventHandler)

}
