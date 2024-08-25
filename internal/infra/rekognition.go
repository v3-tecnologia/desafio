package infra

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

type Rekognition = *rekognition.Client

func NewRekognition() Rekognition {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		_ = fmt.Errorf("unable to load SDK config, %v", err)

		panic("unable to load SDK config, " + err.Error())
	}

	client := rekognition.NewFromConfig(cfg)

	return client

}
