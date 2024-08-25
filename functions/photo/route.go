package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevenmiano/v3/internal/infra"
	photoRepository "github.com/kevenmiano/v3/internal/repository/photo"
	rekognitionRepository "github.com/kevenmiano/v3/internal/repository/rekognition"
	"github.com/kevenmiano/v3/internal/repository/s3"
	route "github.com/kevenmiano/v3/internal/route/upload"
	"github.com/kevenmiano/v3/internal/service/photo"
	"github.com/kevenmiano/v3/internal/shared"
	photoUseCase "github.com/kevenmiano/v3/internal/usecase/photo"
)

func main() {

	logger := infra.NewLogger()

	rekognition := infra.NewRekognition()

	newRekognitionRepository := rekognitionRepository.NewRekognitionRepository(rekognition)

	s3Repository := s3.NewS3Repository()

	newPhotoRepository := photoRepository.NewPhotoRepository(s3Repository, newRekognitionRepository)

	createPhotoUseCase := photoUseCase.NewCreatePhotoUseCase(newPhotoRepository)

	findPhotoUseCase := photoUseCase.NewFindPhotoUseCase(newPhotoRepository)

	photoService := photo.NewPhotoService(logger, createPhotoUseCase, findPhotoUseCase)

	handler := route.NewRoute(photoService)

	lambdaHandler := shared.NewLambdaHandler(handler)

	lambda.Start(lambdaHandler)
}
