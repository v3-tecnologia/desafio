package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/test/mocks"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreateGpsService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockTelemetryRepository(control)
	service := NewTelemetryService(repo)
	macAddres := "4A-66-45-37-8D-05"
	deviceID := 1
	gpsDomain := domain.GpsDomain{
		Latitude:       1,
		Longitude:      2,
		MacAddress:     macAddres,
		CollectionDate: time.Now(),
	}

	t.Run("when it doesn't find the device it returns error", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{}, err_rest.NewBadRequestErr("unable get Device from Mac_address"))

		err := service.CreateGpsService(gpsDomain)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("when it finds the device but it gives an error when creating the gps", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddres,
		}, nil)

		repo.EXPECT().CreateGps(domain.GpsDomain{
			Latitude:       1,
			Longitude:      2,
			DeviceID:       deviceID,
			MacAddress:     macAddres,
			CollectionDate: time.Time{},
		}).Return(err_rest.NewInternalServerError("unable create gps"))

		err := service.CreateGpsService(gpsDomain)

		if err == nil {
			t.FailNow()
		}

	})

	t.Run("success when creating gps", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddres,
		}, nil)

		repo.EXPECT().CreateGps(domain.GpsDomain{
			Latitude:       1,
			Longitude:      2,
			DeviceID:       deviceID,
			MacAddress:     macAddres,
			CollectionDate: time.Time{},
		}).Return(nil)

		err := service.CreateGpsService(gpsDomain)

		if err != nil {
			t.FailNow()
		}
	})
}
