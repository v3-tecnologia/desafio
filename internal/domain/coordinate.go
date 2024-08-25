package domain

import "errors"

type Coordinate struct {
	Latitude  float64 `json:"latitude" dynamodbav:"Latitude" validate:"required,latitude"`
	Longitude float64 `json:"longitude" dynamodbav:"Longitude" validate:"required,longitude"`
}

func NewCoordinate(latitude float64, longitude float64) (Coordinate, error) {

	return Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

func (c *Coordinate) Validate() (bool, error) {

	if c.Latitude < -90 || c.Latitude > 90 {
		return false, ErrInvalidLatitudeCoordinate
	}

	if c.Longitude < -180 || c.Longitude > 180 {
		return false, ErrInvalidLongitudeCoordinate
	}

	return true, nil
}

var (
	ErrInvalidLatitudeCoordinate  = errors.New("invalid latitude")
	ErrInvalidLongitudeCoordinate = errors.New("invalid longitude")
)
