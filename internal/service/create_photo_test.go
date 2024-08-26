package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/test/mocks"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreatePhotoService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mocks.NewMockTelemetryRepository(control)
	service := NewTelemetryService(repo)
	macAddress := "4A-66-45-37-8D-05"
	deviceID := 1
	photoDomain := domain.PhotoDomain{
		Url:            "www.teste/teste.com",
		DeviceID:       deviceID,
		MacAddress:     macAddress,
		CollectionDate: time.Time{},
	}

	t.Run("when it doesn't find the device it returns error", func(t *testing.T) {

		repo.EXPECT().FindDeviceByMAC(macAddress).Return(domain.DeviceDomain{},
			err_rest.NewBadRequestErr("unable get device from mac addres"))

		err := service.CreatePhotoService(photoDomain)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("when it finds device but it gives an error when creating photo", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddress).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddress,
		}, nil)

		repo.EXPECT().CreatePhoto(photoDomain).Return(
			err_rest.NewInternalServerError("unable create photo"),
		)

		err := service.CreatePhotoService(photoDomain)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("create photo successful", func(t *testing.T) {

		repo.EXPECT().FindDeviceByMAC(macAddress).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddress,
		}, nil)

		repo.EXPECT().CreatePhoto(photoDomain).Return(nil)

		err := service.CreatePhotoService(photoDomain)

		if err != nil {
			t.FailNow()
		}
	})

}
