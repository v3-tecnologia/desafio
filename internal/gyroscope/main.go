package gyroscope

import (
	"desafio-backend/pkg/errors"
	"io"
)

type UseCases interface {
	ValidateGyroscope(gyroscope Request) errors.ErrorList
	ParseGyroscope(gyroscope io.ReadCloser) (Request, errors.Error)
	SaveGyroscope(gyroscope Request) (Response, errors.Error)
}

type Main struct{}

func NewMain() UseCases {
	return Main{}
}
