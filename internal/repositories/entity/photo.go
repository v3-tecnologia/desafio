package entity

import "time"

type Photos struct {
	ID             int
	Url            string
	DeviceID       int
	CollectionDate *time.Time
}
