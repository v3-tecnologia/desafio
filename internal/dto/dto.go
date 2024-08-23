package dto

import "time"

type CreateGyroscopeInputDTO struct {
	Name       string  `json:"name" binding:"required"`
	Model      string  `json:"model" binding:"required"`
	X          float64 `json:"x" binding:"required"`
	Y          float64 `json:"y" binding:"required"`
	Z          float64 `json:"z" binding:"required"`
	MacAddress string  `json:"mac_address" binding:"required"`
}

type GyroscopeOutputDTO struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	X          float64   `json:"x"`
	Y          float64   `json:"y"`
	Z          float64   `json:"z"`
	Timestamp  time.Time `json:"timestamp"`
	MacAddress string    `json:"mac_address"`
}
