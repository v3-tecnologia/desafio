package service

import (
	"github/desafio/models"
	"github/desafio/service/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service ProcessData
	repo    *mocks.Repository
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.repo = &mocks.Repository{}
	ts.service = NewService(ts.repo)
}

func (ts *ServiceTestSuite) TestProcessGyroscopeData() {
	ts.Run("success: normal conditions", func() {
		data := models.Gyroscope{
			MacAddress: "00:00:00:00:00:00",
			X:          123.1,
			Y:          213.2,
			Z:          32.3,
			Timestamp:  1724855500}
		ts.repo.On("InsertGyroscopeData", data).Return(nil)
		err := ts.service.ProcessGyroscopeData(data)
		ts.Nil(err)
	})
}

func (ts *ServiceTestSuite) TestProcessGPSData() {
	ts.Run("success: normal conditions", func() {
		data := models.GPS{
			MacAddress: "00:00:00:00:00:00",
			Latitude:   "-5.088889",
			Longitude:  "-42.801944",
			Timestamp:  1724855500}
		ts.repo.On("InsertGPSData", data).Return(nil)
		err := ts.service.ProcessGPSData(data)
		ts.Nil(err)
	})
}

func (ts *ServiceTestSuite) TestProcessPhotoData() {
	ts.Run("success: normal conditions", func() {
		data := models.Photo{
			MacAddress: "00:00:00:00:00:00",
			Photo:      photo,
			Timestamp:  1724855500}
		ts.repo.On("InsertPhoto", data).Return(nil)
		err := ts.service.ProcessPhoto(data)
		ts.Nil(err)
	})
}
