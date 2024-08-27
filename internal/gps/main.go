package gps

import (
	"desafio-backend/internal/device"
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
	db         *gorm.DB
	deviceMain device.UseCases
}

func NewMain(db *gorm.DB, deviceMain device.UseCases) UseCases {
	return Main{db, deviceMain}
}
