package models

import (
	"errors"
	"fmt"
)

type Gyroscope struct {
	*DeviceData `json:"deviceData"`
	X           *float64 `json:"x"`
	Y           *float64 `json:"y"`
	Z           *float64 `json:"z"`
}

func NewGyroscope(d *DeviceData, x, y, z *float64) (*Gyroscope, error) {
	if d == nil {
		return nil, errors.New("device data cannot be nil")
	}

	if x == nil || y == nil || z == nil {
		var nilVars []string
		if x == nil {
			nilVars = append(nilVars, "X")
		}
		if y == nil {
			nilVars = append(nilVars, "Y")
		}
		if z == nil {
			nilVars = append(nilVars, "Z")
		}
		return nil, fmt.Errorf("%v cannot be nil", nilVars)
	}

	return &Gyroscope{
		DeviceData: d,
		X:          x,
		Y:          y,
		Z:          z,
	}, nil
}
