package entity

import (
	"errors"
	"github.com/google/uuid"
)

type GPS struct {
	id        string
	latitude  float64
	longitude float64
	Device
}

func (g *GPS) IsValid() error {
	if g.id == "" {
		return errors.New("invalid id")
	}
	if g.latitude < -90 || g.latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}
	if g.longitude < -180 || g.longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}
	return nil
}

func NewGPS(latitude, longitude float64, macAddress string) (*GPS, error) {
	device, err := NewDevice(macAddress)
	if err != nil {
		return nil, err
	}
	gps := &GPS{
		id:        uuid.New().String(),
		latitude:  latitude,
		longitude: longitude,
		Device:    *device,
	}
	err = gps.IsValid()
	if err != nil {
		return nil, err
	}
	return gps, nil
}

func (g *GPS) GetID() string {
	return g.id
}

func (g *GPS) GetLatitude() float64 {
	return g.latitude
}

func (g *GPS) GetLongitude() float64 {
	return g.longitude
}
