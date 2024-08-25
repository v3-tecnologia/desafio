package models

import "errors"

type Gyroscope struct {
	*DeviceData `json:"deviceData"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           float64 `json:"z"`
}

func NewGyroscope(d *DeviceData, x, y, z float64) (*Gyroscope, error) {
	if d == nil {
		return nil, errors.New("device data cannot be nil")
	}

	return &Gyroscope{
		DeviceData: d,
		X:          x,
		Y:          y,
		Z:          z,
	}, nil
}
