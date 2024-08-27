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
	ers := errors.NewErrorList()

	if gps.MacAddress == "" {
		err := errors.NewError("Missing MacAddress", "MacAddress is required").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGps.CheckMacAddress")
		ers.Append(err)
	}

	macErr := util.IsValidateMacAddress(gps.MacAddress)
	if macErr != nil {
		ers.Append(errors.NewError("Invalid MacAddress format", "MacAddress is not valid").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGps.MacAddressFormat"))
	}

	if gps.Latitude == 0.0 {
		err := errors.NewError("Missing Latitude", "Latitude is required").
			WithMeta("field", "latitude").
			WithOperations("ValidateGps.CheckLatitude")
		ers.Append(err)
	}

	if gps.Longitude == 0.0 {
		err := errors.NewError("Missing Longitude", "Longitude is required").
			WithMeta("field", "longitude").
			WithOperations("ValidateGps.CheckLongitude")
		ers.Append(err)
	}

	if gps.Timestamp.IsZero() {
		err := errors.NewError("Missing Timestamp", "Timestamp is required").
			WithMeta("field", "timestamp").
			WithOperations("ValidateGps.CheckTimestamp")
		ers.Append(err)
	}

	return ers
}

func (entity Request) toResponse() Response {
	return Response{
		MacAddress: entity.MacAddress,
		Timestamp:  entity.Timestamp,
		Latitude:   entity.Latitude,
		Longitude:  entity.Longitude,
	}
}
