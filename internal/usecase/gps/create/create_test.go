package create_test

import (
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/gps"
	"github.com/kevenmiano/v3/internal/usecase/gps"
	"testing"

	"github.com/kevenmiano/v3/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateGpsUseCase_Execute(t *testing.T) {

	gpsRepository := new(r.Repository)

	gpsUseCase := gps.NewCreateGpsUseCase(gpsRepository)

	newGps := dummy.Gps()

	gpsRepository.On("Create", newGps).Return(newGps, nil)

	gpsCreated, err := gpsUseCase.Execute(newGps)

	assert.Equal(t, newGps.ID, gpsCreated.ID)

	assert.Equal(t, newGps.Timestamp, gpsCreated.Timestamp)

	assert.Equal(t, newGps.Coordinate, gpsCreated.Coordinate)

	assert.Equal(t, newGps, gpsCreated)

	assert.Nil(t, err)

	gpsRepository.AssertExpectations(t)

}

func TestInvalidLatitude_Error(t *testing.T) {

	gpsRepository := new(r.Repository)

	gpsUseCase := gps.NewCreateGpsUseCase(gpsRepository)

	newGps := dummy.Gps()

	newGps.Latitude = 91

	expectedErr := domain.ErrInvalidLatitudeCoordinate

	gpsCreated, err := gpsUseCase.Execute(newGps)

	assert.Nil(t, gpsCreated)

	assert.Equal(t, expectedErr, err)

	gpsRepository.AssertExpectations(t)

}

func TestInvalidLongitude_Error(t *testing.T) {

	gpsRepository := new(r.Repository)

	gpsUseCase := gps.NewCreateGpsUseCase(gpsRepository)

	newGps := dummy.Gps()

	newGps.Longitude = 181

	expectedErr := domain.ErrInvalidLongitudeCoordinate

	gpsCreated, err := gpsUseCase.Execute(newGps)

	assert.Nil(t, gpsCreated)

	assert.Equal(t, expectedErr, err)

	gpsRepository.AssertExpectations(t)
}

func TestInvalidDeviceID_Error(t *testing.T) {

	gpsRepository := new(r.Repository)

	gpsUseCase := gps.NewCreateGpsUseCase(gpsRepository)

	newGps := dummy.Gps()

	newGps.DeviceID = ""

	expectedErr := domain.ErrDeviceIdNotFoundGps

	gpsCreated, err := gpsUseCase.Execute(newGps)

	assert.Nil(t, gpsCreated)

	assert.Equal(t, expectedErr, err)

	gpsRepository.AssertExpectations(t)

}

func TestInvalidTimestamp_Error(t *testing.T) {

	gpsRepository := new(r.Repository)

	gpsUseCase := gps.NewCreateGpsUseCase(gpsRepository)

	newGps := dummy.Gps()

	newGps.Timestamp = 0

	expectedErr := domain.ErrTimestampNotFoundGps

	gpsCreated, err := gpsUseCase.Execute(newGps)

	assert.Nil(t, gpsCreated)

	assert.Equal(t, expectedErr, err)

	gpsRepository.AssertExpectations(t)

}
