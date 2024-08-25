package shared

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Request = events.APIGatewayProxyRequest

type Response = events.APIGatewayProxyResponse

type event = events.S3Event

type Handler interface {
	Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type Event interface {
	Handler(ctx context.Context, event events.S3Event) (events.S3Event, error)
}

func NewLambdaHandler(h Handler) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request Request) (response Response, err error) {
		return h.Handler(ctx, request)
	}
}

func NewEventLambdaHandler(h Event) func(ctx context.Context, event events.S3Event) (events.S3Event, error) {
	return func(ctx context.Context, event event) (events.S3Event, error) {
		return h.Handler(ctx, event)
	}
}
