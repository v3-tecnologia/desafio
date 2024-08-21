package usecase

import (
	"bytes"
	"testing"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
	"github.com/charmingruby/g3/internal/telemetry/domain/port"
	"github.com/charmingruby/g3/test/inmemory_adapter"
	"github.com/charmingruby/g3/test/inmemory_repository"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	gpsRepo           *inmemory_repository.GPSInMemoryRepository
	gyroscopeRepo     *inmemory_repository.GyroscopeInMemoryRepository
	photoRepo         *inmemory_repository.PhotoInMemoryRepository
	storageAdapter    *inmemory_adapter.StorageInMemoryAdapter
	recognizerAdapter *inmemory_adapter.RecognizerInMemoryAdapter
	useCase           TelemetryUseCaseRegistry
}

func (s *Suite) SetupSuite() {
	s.gpsRepo = inmemory_repository.NewGPSInMemoryRepository()
	s.gyroscopeRepo = inmemory_repository.NewGyroscopeInMemoryRepository()
	s.photoRepo = inmemory_repository.NewPhotoInMemoryRepository()
	s.storageAdapter = inmemory_adapter.NewStorageInMemoryAdapter()
	s.recognizerAdapter = inmemory_adapter.NewRecognizerInMemoryAdapter()
	s.useCase = NewTelemetryUseCaseRegistry(s.gpsRepo, s.gyroscopeRepo, s.photoRepo, s.storageAdapter, s.recognizerAdapter)
}

func (s *Suite) SetupTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
	s.recognizerAdapter.MockedFaces = []port.DetectedFace{}
	s.storageAdapter.Files = make(map[string]*bytes.Buffer)
}

func (s *Suite) TearDownTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
	s.recognizerAdapter.MockedFaces = []port.DetectedFace{}
	s.storageAdapter.Files = make(map[string]*bytes.Buffer)
}

func (s *Suite) SetupSubTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
	s.recognizerAdapter.MockedFaces = []port.DetectedFace{}
	s.storageAdapter.Files = make(map[string]*bytes.Buffer)
}

func (s *Suite) TearDownSubTest() {
	s.gpsRepo.Items = []entity.GPS{}
	s.gyroscopeRepo.Items = []entity.Gyroscope{}
	s.photoRepo.Items = []entity.Photo{}
	s.recognizerAdapter.MockedFaces = []port.DetectedFace{}
	s.storageAdapter.Files = make(map[string]*bytes.Buffer)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
