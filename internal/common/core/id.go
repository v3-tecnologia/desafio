package core

import "github.com/oklog/ulid/v2"

func NewID() string {
	return ulid.Make().String()
}
