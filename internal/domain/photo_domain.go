package domain

import "time"

type PhotoDomain struct {
	ID             int
	Url            string
	DeviceID       int
	MacAddress     string
	CollectionDate time.Time
}
