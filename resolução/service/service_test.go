package service

import (
	"desafio/models"
	"desafio/service/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service IService
	repo    *mocks.IRepository
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (ts *ServiceTestSuite) SetupTest() {
	ts.repo = &mocks.IRepository{}
	ts.service = NewService(ts.repo)
}

func (ts *ServiceTestSuite) TestProcessGyroscopeData() {
	ts.Run("success: normal conditions", func() {
		data := models.GyroscopeRequest{
			Mac:           "00-00-00-00-00-00",
			X:             111.1,
			Y:             222.2,
			Z:             333.3,
			UnixtimeStamp: 1724603773}
		ts.repo.On("InsertGyroscopeData", data).Return(nil)
		err := ts.service.ProcessGyroscopeData(data)
		ts.Nil(err)
	})
}

func (ts *ServiceTestSuite) TestProcessGpsData() {
	ts.Run("success: normal conditions", func() {
		data := models.GpsRequest{
			Mac:           "00-00-00-00-00-00",
			Lat:           "-18.909762",
			Lon:           "-48.232750",
			UnixtimeStamp: 1724603773}
		ts.repo.On("InsertGpsData", data).Return(nil)
		err := ts.service.ProcessGpsData(data)
		ts.Nil(err)
	})
}

func (ts *ServiceTestSuite) TestProcessPhotoData() {
	ts.Run("success: normal conditions", func() {
		data := models.PhotoRequest{
			Mac:           "00-00-00-00-00-00",
			ImageBase64:   image,
			UnixtimeStamp: 1724603773}
		ts.repo.On("InsertPhotoData", data).Return(nil)
		err := ts.service.ProcessPhotoData(data)
		ts.Nil(err)
	})
}
