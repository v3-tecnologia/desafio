package usecase

import (
	"testing"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
	"github.com/charmingruby/g3/test/inmemory_repository"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	gpsRepo       *inmemory_repository.GPSInMemoryRepository
	gyroscopeRepo *inmemory_repository.GyroscopeInMemoryRepository
	photoRepo     *inmemory_repository.PhotoInMemoryRepository
	useCase       TelemetryUseCaseRegistry
}

func (s *Suite) SetupSuite() {
	s.gpsRepo = inmemory_repository.NewGPSInMemoryRepository()
	s.gyroscopeRepo = inmemory_repository.NewGyroscopeInMemoryRepository()
	s.photoRepo = inmemory_repository.NewPhotoInMemoryRepository()
	s.useCase = NewTelemetryUseCaseRegistry(s.gpsRepo, s.gyroscopeRepo, s.photoRepo)
}

func (s *Suite) SetupTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
}

func (s *Suite) TearDownTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
}

func (s *Suite) SetupSubTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
}

func (s *Suite) TearDownSubTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
