package main

import (
	"log"
	"math"
	"math/rand/v2"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestGPSValidJson(t *testing.T) {
	var g gps
	obj := &g

	valid_json := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"latitude": 3.1415926,
		"longitude": 2.7182818
	}`)

	success := obj.decode(valid_json)
	expected_result := gps{deviceID: "Device1", timestamp: 420, latitude: 3.1415926, longitude: 2.7182818}
	if !success || g != expected_result {
		log.Println(success, g, expected_result)
		t.Fatal("TestGPS failed on valid input")
	}
}

func TestGPSInvalidID(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_value := []byte(`{
		"deviceID": 35,
		"timestamp": 420,
		"latitude": 3.1415926,
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_value)
	if success {
		t.Fatal("TestGPS failed on invalid id")
	}
}

func TestGPSInvalidTime(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_value2 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420.5,
		"latitude": 3.1415926,
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_value2)
	if success {
		t.Fatal("TestGPS failed on invalid timestamp")
	}
}

func TestGPSInvalidlatitude(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_value3 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"latitude": "3.1415926",
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_value3)
	if success {
		t.Fatal("TestGPS failed on invalid latitude")
	}
}

func TestGPSInvalidlongitude(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_value4 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"latitude": 3.1415926,
		"longitude": "2.7182818"
	}`)

	success := obj.decode(invalid_json_value4)
	if success {
		t.Fatal("TestGPS failed on invalid longitude")
	}
}

func TestGPSMissingID(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_id := []byte(`{
		"timestamp": 420,
		"latitude": 3.1415926,
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_id)
	if success {
		t.Fatal("TestGPS.g failed on missing id")
	}
}

func TestGPSMissingTime(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_time := []byte(`{
		"deviceID": "Device1",
		"latitude": 3.1415926,
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_time)
	if success {
		t.Fatal("TestGPS.g failed on missing id")
	}
}

func TestGPSMissinglatitude(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_latitude := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"longitude": 2.7182818
	}`)

	success := obj.decode(invalid_json_latitude)
	if success {
		t.Fatal("TestGPS.g failed on missing id")
	}
}

func TestGPSMissinglongitude(t *testing.T) {
	var g gps
	obj := &g

	invalid_json_longitude := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"latitude": 3.1415926
	}`)

	success := obj.decode(invalid_json_longitude)
	if success {
		t.Fatal("TestGPS.g failed on missing id")
	}
}

func TestGPSPersist(t *testing.T) {
	roundFloat := func(val float64, precision uint) float64 {
		ratio := math.Pow(10, float64(precision))
		return math.Round(val*ratio) / ratio
	}

	g := &gps{roundFloat(rand.Float64(), 5), roundFloat(rand.Float64(), 5), rand.Uint64(), "Device"}

	cfg := mysql.Config{
		User:   "tester",
		Passwd: "test_passwd",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "db_test",
	}

	db, err := connectDatabase(cfg)
	if err != nil {
		log.Println(err)
		t.Fatal("TestGPS failed connecting to database")
	}

	db.Exec("DROP TABLE IF EXISTS gps")
	_, err = db.Exec("CREATE TABLE gps (deviceID varchar(255), time BIGINT unsigned, latitude float(24), longitude float(24))")
	if err != nil {
		log.Println(err)
		t.Fatal("TestGPS failed on creating table")
	}

	err = g.persist(db)
	if err != nil {
		log.Println(g, err)
		t.Fatal("TestGPS failed on persist")
	}

	var obj gps
	result := db.QueryRow("SELECT * FROM gps")
	db.Exec("DROP TABLE gps")
	db.Close()
	if er := result.Scan(&obj.deviceID, &obj.timestamp, &obj.latitude, &obj.longitude); er != nil {
		log.Println(er)
		t.Fatal("TestGPS failed retrieving data")
	}

	if obj != *g {
		log.Println(*g)
		log.Println(obj)
		t.Fatal("Retrieved incorrect value")
	}
}
