package device

import (
	"desafio-backend/pkg/errors"
	"gorm.io/gorm"
)

type UseCases interface {
	FindByMacAddress(macAddress string) (*Device, errors.Error)
	SaveDevice(device Device) (Device, errors.Error)
}

type Main struct {
	db *gorm.DB
}

func NewMain(db *gorm.DB) UseCases {
	return Main{db}
}
