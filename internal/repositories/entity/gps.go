package entity

import "time"

type Gps struct {
	ID             int
	Latitude       float64
	Longitude      float64
	DeviceID       int
	CollectionDate *time.Time
}
