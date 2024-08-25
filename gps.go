package main

import (
	"database/sql"
	"encoding/json"
)

type gps struct {
	latitude, longitude float64
	timestamp           uint64
	deviceID            string
}

func (g *gps) decode(data []byte) bool {
	var m map[string]interface{}
	var valid_device, valid_latitude, valid_longitude, valid_time bool

	err := json.Unmarshal(data, &m)
	if err != nil {
		return false
	}

	g.deviceID, valid_device = validateDevice(m)
	g.timestamp, valid_time = validateTimestamp(m)
	g.latitude, valid_latitude = validateLatitude(m)
	g.longitude, valid_longitude = validateLongitude(m)

	return valid_device && valid_time && valid_latitude && valid_longitude
}

func validateLatitude(vals map[string]interface{}) (float64, bool) {
	latitude, valid := vals["latitude"].(float64)

	if valid {
		valid = (latitude > -90.0 && latitude < 90.0)
	}

	return latitude, valid
}

func validateLongitude(vals map[string]interface{}) (float64, bool) {
	longitude, valid := vals["longitude"].(float64)

	if valid {
		valid = (longitude > -180.0 && longitude < 180.0)
	}

	return longitude, valid
}

func (g *gps) persist(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO gps (deviceID, latitude, longitude, time) VALUES (?, ?, ?, ?)", g.deviceID, g.latitude, g.longitude, g.timestamp)
	return err
}
