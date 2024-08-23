package main

import (
	"database/sql"
	"encoding/json"
)

type gyroscope struct {
	x, y, z   float64
	timestamp uint64
	deviceID  string
}

func (g *gyroscope) decode(data []byte) bool {
	var m map[string]interface{}
	var valid_device, valid_x, valid_y, valid_z bool

	err := json.Unmarshal(data, &m)
	if err != nil {
		return false
	}

	g.deviceID, valid_device = m["deviceID"].(string)
	timestamp, valid_time := m["timestamp"].(float64)
	g.x, valid_x = m["x"].(float64)
	g.y, valid_y = m["y"].(float64)
	g.z, valid_z = m["z"].(float64)

	if valid_time {
		valid_time = (timestamp == float64(uint64(timestamp)))
	}

	g.timestamp = uint64(timestamp)

	return valid_device && valid_time && valid_x && valid_y && valid_z
}

func (g *gyroscope) persist(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO v3.gyroscope (deviceID, x, y, z, time) VALUES (?, ?, ?, ?, ?)", g.deviceID, g.x, g.y, g.z, g.timestamp)
	return err
}
