package shared

import (
	"context"
)

type Multipart[T any] func(ctx context.Context, request Request, file T) (Response, error)
