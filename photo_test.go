package main

import (
	"encoding/base64"
	"os"
	"reflect"
	"testing"
)

func TestValidateFormat(t *testing.T) {
	format, valid := validateFormat(map[string]interface{}{"format": "bmp"})
	if !valid || format != "bmp" {
		t.Fatal(valid, format)
	}

	format, valid = validateFormat(map[string]interface{}{"format": "jpg"})
	if !valid || format != "jpg" {
		t.Fatal(valid, format)
	}

	format, valid = validateFormat(map[string]interface{}{"format": "mp3"})
	if valid {
		t.Fatal(valid, format)
	}
}

func TestValidateImage(t *testing.T) {
	file, err := os.Open("./test_files/image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	img_arr := make([]byte, info.Size())
	_, err = file.Read(img_arr)
	if err != nil {
		t.Fatal(err)
	}

	img_str := base64.StdEncoding.EncodeToString(img_arr)
	img, valid := validateImage(map[string]interface{}{"image": img_str})
	if !valid || !reflect.DeepEqual(img, img_arr) {
		t.Fatal(valid, reflect.DeepEqual(img, img_arr))
	}
}

func TestDecodeImage(t *testing.T) {
	file, err := os.Open("./test_files/image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	img_arr := make([]byte, info.Size())
	_, err = file.Read(img_arr)
	if err != nil {
		t.Fatal(err)
	}

	img_str := base64.StdEncoding.EncodeToString(img_arr)
	device_field := "\"deviceID\": \"Device1\""
	time_field := "\"timestamp\": 420"
	image_field := "\"image\": \"" + img_str + "\""
	format_field := "\"format\": \"jpg\""
	valid_input := "{" + device_field + "," + time_field + "," + image_field + "," + format_field + "}"
	invalid_input := "{" + device_field + "," + image_field + "," + format_field + "}"

	json := []byte(valid_input)
	p := &photo{}
	valid := p.decode(json)
	result := photo{img_arr, "./Device1-420.jpg", "jpg", 420, "Device1"}
	same := reflect.DeepEqual(p.image, result.image)
	if !valid || p.deviceID != result.deviceID || p.file != result.file || p.format != result.format || p.timestamp != result.timestamp || !same {
		t.Fatal(valid, *p)
	}

	json = []byte(invalid_input)
	valid = p.decode(json)
	if valid {
		t.Fatal(valid)
	}
}
