package gyroscope

import (
	"desafio-backend/pkg/errors"
	"desafio-backend/pkg/logger"
	"desafio-backend/util"
	"encoding/json"
	"io"
	"time"
)

// Gyroscope limit in  (°/s) - ±2000°/s is an example value for high sensibility gyroscopes
const (
	minValue = -2000.0
	maxValue = 2000.0
)

type Request struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	XAxis      float64   `json:"xAxis"`
	YAxis      float64   `json:"yAxis"`
	ZAxis      float64   `json:"zAxis"`
}

type Response struct {
	MacAddress string    `json:"macAddress"`
	Timestamp  time.Time `json:"timestamp"`
	XAxis      float64   `json:"xAxis"`
	YAxis      float64   `json:"yAxis"`
	ZAxis      float64   `json:"zAxis"`
}

func (main Main) ParseGyroscope(gyroscope io.ReadCloser) (Request, errors.Error) {
	var nGyroscope *Request

	err := json.NewDecoder(gyroscope).Decode(&nGyroscope)
	if err != nil {
		logger.Error(util.GeneralParseError, "ParseGyroscope", err, gyroscope)
		return Request{}, errors.NewError("Connot decode data", err.Error()).
			WithOperations("ParseGyroscope.Decode")
	}
	return *nGyroscope, nil
}

func (main Main) SaveGyroscope(gyroscope Request) (Response, errors.Error) {
	// TODO save the data received into a database
	return gyroscope.toResponse(), nil
}

func (main Main) ValidateGyroscope(gyroscope Request) errors.ErrorList {
	ers := errors.NewErrorList()

	if gyroscope.MacAddress == "" {
		ers.Append(errors.NewError("Missing MacAddress", "MacAddress is required").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGyroscope.MacAddress"))
	}

	macErr := util.IsValidateMacAddress(gyroscope.MacAddress)
	if macErr != nil {
		ers.Append(errors.NewError("Invalid MacAddress format", "MacAddress is not valid").
			WithMeta("field", "macAddress").
			WithOperations("ValidateGyroscope.MacAddressFormat"))
	}

	if gyroscope.Timestamp.IsZero() {
		ers.Append(errors.NewError("Missing Timestamp", "Timestamp is required").
			WithMeta("field", "timestamp").
			WithOperations("ValidateGyroscope.Timestamp"))
	}

	if gyroscope.XAxis == 0 && gyroscope.YAxis == 0 && gyroscope.ZAxis == 0 {
		ers.Append(errors.NewError("Invalid Gyroscope Data", "At least one axis value must be non-zero").
			WithOperations("ValidateGyroscope.GyroscopeData"))
	}

	if gyroscope.XAxis < minValue || gyroscope.XAxis > maxValue {
		ers.Append(errors.NewError("Invalid X axis", "X axis value must be within defined range").
			WithOperations("ValidateGyroscope.XAxis"))
	}

	if gyroscope.YAxis < minValue || gyroscope.YAxis > maxValue {
		ers.Append(errors.NewError("Invalid Y axis", "Y axis value must be within defined range").
			WithOperations("ValidateGyroscope.YAxis"))
	}

	if gyroscope.ZAxis < minValue || gyroscope.ZAxis > maxValue {
		ers.Append(errors.NewError("Invalid Z axis", "Z axis value must be within defined range").
			WithOperations("ValidateGyroscope.ZAxis"))
	}

	return ers
}

func (entity Request) toResponse() Response {
	return Response{
		MacAddress: entity.MacAddress,
		Timestamp:  entity.Timestamp,
		XAxis:      entity.XAxis,
		YAxis:      entity.YAxis,
		ZAxis:      entity.ZAxis,
	}
}
