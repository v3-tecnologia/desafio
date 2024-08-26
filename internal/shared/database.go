package shared

import "context"

type Database[T any] interface {
	Put(ctx context.Context, item T) error
}
