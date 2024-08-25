package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestConnectDatabase(t *testing.T) {
	correct_cfg := mysql.Config{
		User:   "tester",
		Passwd: "test_passwd",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "db_test",
	}

	db, err := connectDatabase("mysql", correct_cfg.FormatDSN())
	if db == nil || err != nil {
		log.Println(db, err)
		t.Fatal("Test failed connecting with correct cfg")
	}
}

func TestConnectDatabaseFail(t *testing.T) {
	correct_cfg := mysql.Config{
		User:   "fake_user",
		Passwd: "fake_passwd",
	}

	db, err := connectDatabase("mysql", correct_cfg.FormatDSN())
	if db != nil || err == nil {
		log.Println(db, err)
		t.Fatal("Test failed connecting with bad cfg")
	}
}

func TestValidateDevice(t *testing.T) {
	valid_result, success := validateDevice(map[string]interface{}{"deviceID": "Valid device"})
	invalid_result, fail := validateDevice(map[string]interface{}{"deviceID": ""})
	if !success || fail || valid_result != "Valid device" || invalid_result != "" {
		t.Fatal(success, valid_result, fail, invalid_result)
	}
}

func TestValidateTimestamp(t *testing.T) {
	valid_result, success := validateTimestamp(map[string]interface{}{"timestamp": 420.0})
	invalid_result, fail := validateTimestamp(map[string]interface{}{"timestamp": 420.3})
	if !success || fail || valid_result != 420 || invalid_result != 420.0 {
		t.Fatal(success, valid_result, fail, invalid_result)
	}
}

type fake_db_table struct {
	counter       uint64
	decodeResult  bool
	persistResult error
}

func (f *fake_db_table) decode(b []byte) bool {
	f.counter++
	return f.decodeResult
}

func (f *fake_db_table) persist(db *sql.DB) error {
	f.counter++
	return f.persistResult
}

type fake_reader struct{}

func (f fake_reader) Read(b []byte) (int, error) { return 0, nil }
func (f fake_reader) Close() error               { return nil }

type fake_response_writer struct {
	code   int
	header http.Header
}

func (f *fake_response_writer) Header() http.Header         { return f.header }
func (f *fake_response_writer) Write(b []byte) (int, error) { return 0, nil }
func (f *fake_response_writer) WriteHeader(i int)           { f.code = i }

func TestProperHandler(t *testing.T) {
	f := &fake_response_writer{header: http.Header{}}
	tst := &fake_db_table{0, true, nil}
	handler := makeHandler(func() db_table { return tst }, nil)
	handler(f, &http.Request{ContentLength: 0, Body: fake_reader{}})
	if tst.counter != 2 || f.code != 0 {
		log.Println(tst.counter, f.code)
		t.Fatal("Handler result mismatch")
	}
}

func TestFailingHandler(t *testing.T) {
	f := &fake_response_writer{header: http.Header{}}
	tst := &fake_db_table{0, false, nil}
	handler := makeHandler(func() db_table { return tst }, nil)
	handler(f, &http.Request{ContentLength: 0, Body: fake_reader{}})
	if tst.counter != 1 || f.code != 400 {
		log.Println(tst.counter, f.code)
		t.Fatal("Handler result mismatch")
	}
}

func TestFailingDatabase(t *testing.T) {
	f := &fake_response_writer{header: http.Header{}}
	tst := &fake_db_table{0, true, errors.New("Persist fail")}
	handler := makeHandler(func() db_table { return tst }, nil)
	handler(f, &http.Request{ContentLength: 0, Body: fake_reader{}})
	if tst.counter != 2 || f.code != 500 {
		log.Println(tst.counter, f.code)
		t.Fatal("Handler result mismatch")
	}
}
