package models

// Modelo utiizado na requisição dos dados de giroscópio
type Gyroscope struct {
	MacAddress string  `json:"macAddress" validate:"mac,required"`
	Timestamp  int     `json:"timestamp" validate:"required,number"`
	X          float64 `json:"x" validate:"required,number"`
	Y          float64 `json:"y" validate:"required,number"`
	Z          float64 `json:"z" validate:"required,number"`
}

// Modelo utiizado na requisição dos dados de GPS
type GPS struct {
	MacAddress string `json:"macAddress" validate:"mac,required"`
	Timestamp  int    `json:"timestamp" validate:"required,number"`
	Latitude   string `json:"latitude" validate:"required,latitude"`
	Longitude  string `json:"longitude" validate:"required,longitude"`
}

// Modelo utiizado na requisição dos dados da foto
type Photo struct {
	MacAddress string `json:"macAddress" validate:"mac,required"`
	Photo      string `json:"photo" validate:"base64,required"`
	Timestamp  int    `json:"timestamp" validate:"required,number"`
}
