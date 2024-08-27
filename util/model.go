package util

import (
	"time"
)

type BaseModel struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}
