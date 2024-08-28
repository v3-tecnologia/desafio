package service_test

import (
	"github/desafio/models"
	"github/desafio/service"
	"github/desafio/service/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Aqui estão as funções para testes unitários nas funções do service
// Foram utilizadas test suite em conjunto com o mockery

type ServiceTestSuite struct {
	suite.Suite
	service service.ProcessData
	repo    *mocks.Repository
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.repo = &mocks.Repository{}
	ts.service = service.NewService(ts.repo)
}

func (ts *ServiceTestSuite) TestProcessGyroscopeData() {
	ts.Run("Success: all data", func() {
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
	ts.Run("Success: all data", func() {
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
	ts.Run("Success: all data", func() {
		data := models.Photo{
			MacAddress: "00:00:00:00:00:00",
			Photo:      service.Photo,
			Timestamp:  1724855500}
		ts.repo.On("InsertPhoto", data).Return(nil)
		err := ts.service.ProcessPhoto(data)
		ts.Nil(err)
	})
}
