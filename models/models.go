package models

type Gyroscope struct {
	MacAddress string  `json:"macAddress" validate:"mac,required"`
	Timestamp  int     `json:"timestamp" validate:"required,number"`
	X          float64 `json:"x" validate:"required"`
	Y          float64 `json:"y" validate:"required"`
	Z          float64 `json:"z" validate:"required"`
}

type GPS struct {
	MacAddress string `json:"macAddress" validate:"mac,required"`
	Timestamp  int    `json:"timestamp" validate:"required,number"`
	Latitude   string `json:"latitude" validate:"required"`
	Longitude  string `json:"longitude" validate:"required"`
}

type Photo struct {
	MacAddress string `json:"macAddress" validate:"mac,required"`
	Photo      string `json:"photo" validate:"base64,required"`
	Timestamp  int    `json:"timestamp" validate:"required,number"`
}
