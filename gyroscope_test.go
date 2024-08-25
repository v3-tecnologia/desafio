package main

import (
	"log"
	"math"
	"math/rand/v2"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestGyroscopeValidJson(t *testing.T) {
	var g gyroscope
	obj := &g

	valid_json := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(valid_json)
	expected_result := gyroscope{deviceID: "Device1", timestamp: 420, x: 3.1415926, y: 2.7182818, z: 1.4142135}
	if !success || g != expected_result {
		t.Fatal("TestGyroscope failed on valid input")
	}
}

func TestGyroscopeInvalidID(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_value := []byte(`{
		"deviceID": 35,
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_value)
	if success {
		t.Fatal("TestGyroscope failed on invalid id")
	}
}

func TestGyroscopeInvalidTime(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_value2 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420.5,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_value2)
	if success {
		t.Fatal("TestGyroscope failed on invalid timestamp")
	}
}

func TestGyroscopeInvalidx(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_value3 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": "3.1415926",
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_value3)
	if success {
		t.Fatal("TestGyroscope failed on invalid x")
	}
}

func TestGyroscopeInvalidY(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_value4 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": "2.7182818",
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_value4)
	if success {
		t.Fatal("TestGyroscope failed on invalid y")
	}
}

func TestGyroscopeInvalidZ(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_value5 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": "1.4142135"
	}`)

	success := obj.decode(invalid_json_value5)
	if success {
		t.Fatal("TestGyroscope failed on invalid z")
	}
}

func TestGyroscopeMissingID(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_id := []byte(`{
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_id)
	if success {
		t.Fatal("TestGyroscope failed on missing id")
	}
}

func TestGyroscopeMissingTime(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_time := []byte(`{
		"deviceID": "Device1",
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_time)
	if success {
		t.Fatal("TestGyroscope failed on missing id")
	}
}

func TestGyroscopeMissingX(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_x := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_x)
	if success {
		t.Fatal("TestGyroscope failed on missing id")
	}
}

func TestGyroscopeMissingY(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_y := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"z": 1.4142135
	}`)

	success := obj.decode(invalid_json_y)
	if success {
		t.Fatal("TestGyroscope failed on missing id")
	}
}

func TestGyroscopeMissingZ(t *testing.T) {
	var g gyroscope
	obj := &g

	invalid_json_z := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818
	}`)

	success := obj.decode(invalid_json_z)
	if success {
		t.Fatal("TestGyroscope failed on missing id")
	}
}

func TestGyroscopePersist(t *testing.T) {
	roundFloat := func(val float64, precision uint) float64 {
		ratio := math.Pow(10, float64(precision))
		return math.Round(val*ratio) / ratio
	}

	g := &gyroscope{roundFloat(rand.Float64(), 5), roundFloat(rand.Float64(), 5), roundFloat(rand.Float64(), 5), rand.Uint64(), "Device"}

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
		t.Fatal("TestGyroscope failed connecting to database")
	}

	db.Exec("DROP TABLE IF EXISTS gyroscope")
	_, err = db.Exec("CREATE TABLE gyroscope (deviceID varchar(255), time BIGINT unsigned, x float(24), y float(24), z float(24))")
	if err != nil {
		log.Println(err)
		t.Fatal("TestGyroscope failed on creating table")
	}

	err = g.persist(db)
	if err != nil {
		log.Println(g, err)
		t.Fatal("TestGyroscope failed on persist")
	}

	var obj gyroscope
	result := db.QueryRow("SELECT * FROM gyroscope")
	db.Exec("DROP TABLE gyroscope")
	db.Close()
	if er := result.Scan(&obj.deviceID, &obj.timestamp, &obj.x, &obj.y, &obj.z); er != nil {
		log.Println(er)
		t.Fatal("TestGyroscope failed retrieving data")
	}

	if obj != *g {
		log.Println(*g)
		log.Println(obj)
		t.Fatal("Retrieved incorrect value")
	}
}
