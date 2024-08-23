package main

import (
	"database/sql"
	"encoding/json"
	"log"
)

type gyroscope struct {
	x, y, z   float64
	timestamp uint64
	deviceID  string
}

func (g *gyroscope) decode(data []byte) bool {
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	deviceID, valid_device := m["deviceID"].(string)
	timestamp := m["timestamp"].(float64)
	valid_time := timestamp == float64(uint64(timestamp))
	x, valid_x := m["x"].(float64)
	y, valid_y := m["y"].(float64)
	z, valid_z := m["z"].(float64)

	g.deviceID = deviceID
	g.timestamp = uint64(timestamp)
	g.x = x
	g.y = y
	g.z = z

	return valid_device && valid_time && valid_x && valid_y && valid_z
}

func (g *gyroscope) persist(db *sql.DB) bool {
	result, err := db.Exec("INSERT INTO v3.gyroscope (deviceID, x, y, z, time) VALUES (?, ?, ?, ?, ?)", g.deviceID, g.x, g.y, g.z, g.timestamp)
	log.Println(result, err)
	return err == nil
}
