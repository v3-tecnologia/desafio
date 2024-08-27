package gps

import (
	"desafio-backend/pkg/errors"
	"io"
)

type UseCases interface {
	ValidateGps(gps Request) errors.ErrorList
	ParseGps(gps io.ReadCloser) (Request, errors.Error)
	SaveGps(gps Request) (Response, errors.Error)
}

type Main struct{}

func NewMain() UseCases {
	return Main{}
}
