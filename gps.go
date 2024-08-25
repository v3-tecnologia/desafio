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
	var valid_device, valid_latitude, valid_longitude bool

	err := json.Unmarshal(data, &m)
	if err != nil {
		return false
	}

	g.deviceID, valid_device = m["deviceID"].(string)
	timestamp, valid_time := m["timestamp"].(float64)
	g.latitude, valid_latitude = m["latitude"].(float64)
	g.longitude, valid_longitude = m["longitude"].(float64)

	if valid_time {
		valid_time = (timestamp == float64(uint64(timestamp)))
	}

	g.timestamp = uint64(timestamp)

	return valid_device && valid_time && valid_latitude && valid_longitude
}

func (g *gps) persist(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO gps (deviceID, latitude, longitude, time) VALUES (?, ?, ?, ?)", g.deviceID, g.latitude, g.longitude, g.timestamp)
	return err
}
