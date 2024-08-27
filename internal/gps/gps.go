package gps

import (
	"desafio-backend/pkg/errors"
	"desafio-backend/pkg/logger"
	"desafio-backend/util"
	"encoding/json"
	"io"
	"time"
)

type Request struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}

type Response struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}

func (main Main) ParseGps(gps io.ReadCloser) (Request, errors.Error) {
	var nGps *Request

	err := json.NewDecoder(gps).Decode(&nGps)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseGps", err, gps)
		return Request{}, errors.NewError("Connot decode data", err.Error()).
			WithOperations("ParseGps.Decode")
	}
	return *nGps, nil
}

func (main Main) SaveGps(gps Request) (Response, errors.Error) {
	// TODO save the data received into a database
	return gps.toResponse(), nil
}

func (main Main) ValidateGps(gps Request) errors.ErrorList {
	// TODO validate data received
	return nil
}

func (entity Request) toResponse() Response {
	return Response{
		MacAddress: entity.MacAddress,
		Timestamp:  entity.Timestamp,
		Latitude:   entity.Latitude,
		Longitude:  entity.Longitude,
	}
}
