package models

import "github.com/go-playground/validator/v10"

type Gyroscope struct {
	MacAddress string  `json:"macaddress" validate:"required"`
	Timestamp  int     `json:"timestamp" validate:"required"`
	X          float64 `json:"x" validate:"required"`
	Y          float64 `json:"y" validate:"required"`
	Z          float64 `json:"z" validate:"required"`
}

type GPS struct {
	MacAddress string  `json:"macaddress" validate:"required"`
	Timestamp  int     `json:"timestamp" validate:"required"`
	Latitude   string `json:"latitude" validate:"required"`
	Longitude  string `json:"longitude" validate:"required"`
}

type Photo struct {
	MacAddress string `json:"macaddress" validate:"required"`
	Photo      string `json:"photo" validate:"base64,required"`
	Timestamp  int    `json:"timestamp" validate:"required"`
}

func ValidateGyroscopeData(gyroscope *Gyroscope) error {
	validate := validator.New()
	err := validate.Struct(gyroscope)
	if err != nil {
		return err
	}

	return nil
}

func ValidateGPSData(gps *GPS) error {
	validate := validator.New()
	err := validate.Struct(gps)
	if err != nil {
		return err
	}

	return nil
}

func ValidatePhotoData(photo *Photo) error {
	validate := validator.New()
	err := validate.Struct(photo)
	if err != nil {
		return err
	}

	return nil
}