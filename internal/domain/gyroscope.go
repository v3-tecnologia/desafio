package domain

import (
	"errors"
	"github.com/kevenmiano/v3/internal/adapter/date"
	"github.com/kevenmiano/v3/internal/adapter/uuid"
)

type Gyroscope struct {
	ID        string  `json:"id" dynamodbav:"ID"`
	DeviceID  string  `json:"deviceId" dynamodbav:"DeviceID" validate:"required"`
	Timestamp int64   `json:"timestamp" dynamodbav:"Timestamp" validate:"required"`
	X         float64 `json:"x" dynamodbav:"X" validate:"required"`
	Y         float64 `json:"y" dynamodbav:"Y" validate:"required"`
	Z         float64 `json:"z" dynamodbav:"Z" validate:"required"`
}

type GyroscopeDto struct {
	ID        string  `json:"id"`
	DeviceID  string  `json:"deviceId"`
	Timestamp int64   `json:"timestamp"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

func NewGyroscope(d *GyroscopeDto) *Gyroscope {

	id := uuid.NewUUIDAdapter()

	timestamp := date.NewDateAdapter()

	return &Gyroscope{
		ID:        id.Value(),
		DeviceID:  d.DeviceID,
		Timestamp: timestamp.Value(),
		X:         d.X,
		Y:         d.Y,
		Z:         d.Z,
	}
}

func (g *Gyroscope) SetDeviceID(deviceID string) {
	g.DeviceID = deviceID
}

func (g *Gyroscope) Validate() (bool, error) {

	if g.DeviceID == "" {
		return false, ErrDeviceIDGyroscope
	}
	if g.Timestamp == 0 {
		return false, ErrTimestampGyroscope
	}
	return true, nil
}

var (
	ErrDeviceIDGyroscope  = errors.New("DeviceID not found")
	ErrTimestampGyroscope = errors.New("timestamp not found")
)
