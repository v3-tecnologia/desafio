package infra

import (
	"context"
	"github.com/kevenmiano/v3/internal/domain"
	"os"
)

func NewGpsDatabase() *Dynamo[domain.GPS] {
	dynamoClient := NewDynamoClient(context.Background())
	return NewDynamo[domain.GPS](dynamoClient, os.Getenv("LOCATION_TABLE"))
}

func NewGyroscopeDatabase() *Dynamo[domain.Gyroscope] {
	dynamoClient := NewDynamoClient(context.Background())
	return NewDynamo[domain.Gyroscope](dynamoClient, os.Getenv("LOCATION_TABLE"))

}
