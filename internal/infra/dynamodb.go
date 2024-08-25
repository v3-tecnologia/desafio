package infra

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDB struct {
	client *dynamodb.Client
}

type Dynamo[T any] struct {
	database  *DynamoDB
	tableName string
}

func NewDynamoClient(ctx context.Context) *DynamoDB {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDB{
		client: client,
	}
}

func NewDynamo[T any](database *DynamoDB, tableName string) *Dynamo[T] {
	return &Dynamo[T]{
		database:  database,
		tableName: tableName,
	}
}

func (s *Dynamo[T]) Put(ctx context.Context, item *T) error {
	av, err := attributevalue.MarshalMap(item)

	if err != nil {
		return fmt.Errorf("unable to marshal item: %v", err)
	}

	_, err = s.database.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &s.tableName,
		Item:      av,
	})

	if err != nil {
		return fmt.Errorf("unable to put item: %v", err)
	}

	return nil
}
