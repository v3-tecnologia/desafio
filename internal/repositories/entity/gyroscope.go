package entity

import "time"

type Gyroscopes struct {
	ID             int
	X              float64
	Y              float64
	Z              float64
	DeviceID       int
	CollectionDate *time.Time
}
