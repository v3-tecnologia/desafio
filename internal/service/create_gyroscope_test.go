package service

import (
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/test/mocks"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreateGyroscope(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mocks.NewMockTelemetryRepository(control)
	service := NewTelemetryService(repo)
	deviceID := 1
	macAddres := "4A-66-45-37-8D-05"
	gyroscopeDomain := domain.GyroscopeDomain{
		X:              10,
		Y:              10,
		Z:              10,
		DeviceID:       deviceID,
		MacAddress:     macAddres,
		CollectionDate: time.Time{},
	}

	t.Run("when it doesn't find the device it returns error", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{},
			err_rest.NewBadRequestErr("unable get device from mac addres"))

		err := service.CreateGyroscopeService(gyroscopeDomain)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("when it finds device but it gives an error when creating gyroscope", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddres,
		}, nil)

		repo.EXPECT().CreateGyroscope(gyroscopeDomain).Return(
			err_rest.NewInternalServerError("unable create gyroscope"),
		)

		err := service.CreateGyroscopeService(gyroscopeDomain)

		if err == nil {
			t.FailNow()
		}
	})

	t.Run("create gyroscope successful", func(t *testing.T) {
		repo.EXPECT().FindDeviceByMAC(macAddres).Return(domain.DeviceDomain{
			ID:  1,
			Mac: macAddres,
		}, nil)

		repo.EXPECT().CreateGyroscope(gyroscopeDomain).Return(nil)

		err := service.CreateGyroscopeService(gyroscopeDomain)

		if err != nil {
			t.FailNow()
		}
	})
}
