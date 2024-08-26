package models

type GyroscopeRequest struct {
	Mac           string  `json:"macAddr" validate:"required,mac"`
	X             float64 `json:"x" validate:"required,number"`
	Y             float64 `json:"y" validate:"required,number"`
	Z             float64 `json:"z" validate:"required,number"`
	UnixtimeStamp int64   `json:"timeStamp" validate:"required,number"`
}

type GpsRequest struct {
	Mac           string `json:"macAddr" validate:"required,mac"`
	Lat           string `json:"latitude" validate:"required,latitude"`
	Lon           string `json:"longitude" validate:"required,longitude"`
	UnixtimeStamp int64  `json:"timeStamp" validate:"required,number"`
}

type PhotoRequest struct {
	Mac           string `json:"macAddr" validate:"required,mac"`
	ImageBase64   string `json:"image" validate:"required"`
	UnixtimeStamp int64  `json:"timeStamp" validate:"required,number"`
}
