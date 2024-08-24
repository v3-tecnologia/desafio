package models

import "errors"

type GPS struct {
	*DeviceData
	Latitude, Longitude float64
}

func NewGPS(d *DeviceData, la, lo float64) (*GPS, error) {
	if d == nil {
		return nil, errors.New("device cannot be nil")
	}
	if la < -90.0 || la > 90.0 {
		return nil, errors.New("latitude must be between -90 and 90")
	}
	if lo < -180.0 || lo > 180.0 {
		return nil, errors.New("longitude must be between -180 and 180")
	}
	return &GPS{
		DeviceData: d,
		Latitude:   la,
		Longitude:  lo,
	}, nil
}
