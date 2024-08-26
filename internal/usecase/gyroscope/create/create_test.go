package create_test

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/gyroscope"
	"github.com/kevenmiano/v3/internal/usecase/gyroscope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGyroscopeUseCase_Execute(t *testing.T) {

	gyroscopeRepository := new(r.Repository)

	gyroscopeUseCase := gyroscope.NewCreateGyroscopeUseCase(gyroscopeRepository)

	newGyroscope := dummy.Gyroscope()

	gyroscopeRepository.On("Create", newGyroscope).Return(newGyroscope, nil)

	gyroscopeCreated, err := gyroscopeUseCase.Execute(newGyroscope)

	assert.Equal(t, gyroscopeCreated, newGyroscope)

	assert.Nil(t, err)

	gyroscopeRepository.AssertExpectations(t)

}

func TestInvalidDeviceID_Error(t *testing.T) {

	gyroscopeRepository := new(r.Repository)

	gyroscopeUseCase := gyroscope.NewCreateGyroscopeUseCase(gyroscopeRepository)

	newGyroscope := dummy.Gyroscope()

	newGyroscope.DeviceID = ""

	expectedErr := domain.ErrDeviceIdNotFoundGps

	gyroscopeCreated, err := gyroscopeUseCase.Execute(newGyroscope)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, gyroscopeCreated)

	gyroscopeRepository.AssertExpectations(t)

}

func TestInvalidTimestamp_Error(t *testing.T) {

	gyroscopeRepository := new(r.Repository)

	gyroscopeUseCase := gyroscope.NewCreateGyroscopeUseCase(gyroscopeRepository)

	newGyroscope := dummy.Gyroscope()

	newGyroscope.Timestamp = 0

	expectedErr := domain.ErrTimestampGyroscope

	gyroscopeCreated, err := gyroscopeUseCase.Execute(newGyroscope)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, gyroscopeCreated)

}
