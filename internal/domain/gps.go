package domain

import (
	"errors"
	"github.com/kevenmiano/v3/internal/adapter/date"
	"github.com/kevenmiano/v3/internal/adapter/uuid"
)

type GPS struct {
	ID        string `json:"id" dynamodbav:"ID"`
	DeviceID  string `json:"deviceId" dynamodbav:"DeviceID" validate:"required"`
	Timestamp int64  `json:"timestamp" dynamodbav:"Timestamp" validate:"required"`
	Coordinate
}

type CoordinateDto struct {
	Latitude  float64 `json:"latitude"  validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type GPSDto struct {
	DeviceID string
	CoordinateDto
}

func NewGPS(d *GPSDto) *GPS {

	id := uuid.NewUUIDAdapter()

	timestamp := date.NewDateAdapter()

	return &GPS{
		ID:        id.Value(),
		DeviceID:  d.DeviceID,
		Timestamp: timestamp.Value(),
		Coordinate: Coordinate{
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
		},
	}
}

func (g *GPS) SetDeviceID(deviceID string) {
	g.DeviceID = deviceID
}

func (g *GPS) Validate() (bool, error) {

	coordinate, _ := NewCoordinate(g.Latitude, g.Longitude)

	if ok, err := coordinate.Validate(); !ok {
		return false, err
	}

	if g.DeviceID == "" {
		return false, ErrDeviceIdNotFoundGps
	}

	if g.Timestamp == 0 {
		return false, ErrTimestampNotFoundGps
	}

	return true, nil
}

var (
	ErrDeviceIdNotFoundGps  = errors.New("DeviceID not found")
	ErrTimestampNotFoundGps = errors.New("timestamp not found")
)
