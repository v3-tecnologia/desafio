package request

import "time"

type GyroscopeRequest struct {
	X              float64   `json:"x"`
	Y              float64   `json:"y"`
	Z              float64   `json:"z"`
	MacAddress     string    `json:"mac_address"`
	CollectionDate time.Time `json:"collection_date"`
}
type GpsRequest struct {
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	MacAddress     string    `json:"mac_address"`
	CollectionDate time.Time `json:"collection_date"`
}
type PhotoRequest struct {
	Url            string    `json:"url"`
	MacAddress     string    `json:"mac_address"`
	CollectionDate time.Time `json:"collection_date"`
}
