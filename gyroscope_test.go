package main

import "testing"

func TestValidJson(t *testing.T) {
	var g gyroscope
	gyro := &g

	valid_json := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(valid_json)
	expected_result := gyroscope{deviceID: "Device1", timestamp: 420, x: 3.1415926, y: 2.7182818, z: 1.4142135}
	if !success || g != expected_result {
		t.Fatal("Test failed on valid input")
	}
}

func TestInvalidID(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_value := []byte(`{
		"deviceID": 35,
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_value)
	if success {
		t.Fatal("Test failed on invalid id")
	}
}

func TestInvalidTime(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_value2 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420.5,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_value2)
	if success {
		t.Fatal("Test failed on invalid timestamp")
	}
}

func TestInvalidx(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_value3 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": "3.1415926",
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_value3)
	if success {
		t.Fatal("Test failed on invalid x")
	}
}

func TestInvalidY(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_value4 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": "2.7182818",
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_value4)
	if success {
		t.Fatal("Test failed on invalid y")
	}
}

func TestInvalidZ(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_value5 := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": "1.4142135"
	}`)

	success := gyro.decode(invalid_json_value5)
	if success {
		t.Fatal("Test failed on invalid z")
	}
}

func TestMissingID(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_id := []byte(`{
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_id)
	if success {
		t.Fatal("Teste failed on missing id")
	}
}

func TestMissingTime(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_time := []byte(`{
		"deviceID": "Device1",
		"x": 3.1415926,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_time)
	if success {
		t.Fatal("Teste failed on missing id")
	}
}

func TestMissingX(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_x := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"y": 2.7182818,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_x)
	if success {
		t.Fatal("Teste failed on missing id")
	}
}

func TestMissingY(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_y := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"z": 1.4142135
	}`)

	success := gyro.decode(invalid_json_y)
	if success {
		t.Fatal("Teste failed on missing id")
	}
}

func TestMissingZ(t *testing.T) {
	var g gyroscope
	gyro := &g

	invalid_json_z := []byte(`{
		"deviceID": "Device1",
		"timestamp": 420,
		"x": 3.1415926,
		"y": 2.7182818
	}`)

	success := gyro.decode(invalid_json_z)
	if success {
		t.Fatal("Teste failed on missing id")
	}
}
