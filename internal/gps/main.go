package gps

import (
	"desafio-backend/pkg/errors"
	"gorm.io/gorm"
	"io"
)

type UseCases interface {
	ValidateGps(gps Request) errors.ErrorList
	ParseGps(gps io.ReadCloser) (Request, errors.Error)
	SaveGps(gps Request) (Response, errors.Error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) UseCases {
	return Main{}
}
