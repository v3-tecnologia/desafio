package models

import (
	"time"
)

type DeviceData struct {
	MAC       string
	Timestamp time.Time
}
