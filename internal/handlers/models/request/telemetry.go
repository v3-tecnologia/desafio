package request

import "time"

type FieldsRequiredTelemetry struct {
	MacAddress     string    `json:"mac_address"`
	CollectionDate time.Time `json:"collection_date"`
}

type GyroscopeRequest struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	FieldsRequiredTelemetry
}
type GpsRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	FieldsRequiredTelemetry
}
type PhotoRequest struct {
	Url string `json:"url"`
	FieldsRequiredTelemetry
}
