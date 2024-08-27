package gyroscope

import (
	"desafio-backend/pkg/errors"
	"gorm.io/gorm"
	"io"
)

type UseCases interface {
	ValidateGyroscope(gyroscope Request) errors.ErrorList
	ParseGyroscope(gyroscope io.ReadCloser) (Request, errors.Error)
	SaveGyroscope(gyroscope Request) (Response, errors.Error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) UseCases {
	return Main{}
}
