package gyroscope

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
	// TODO validate data received
	return nil
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
