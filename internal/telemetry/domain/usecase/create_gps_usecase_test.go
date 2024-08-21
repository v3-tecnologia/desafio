package usecase

import "github.com/charmingruby/g3/internal/telemetry/domain/dto"

func (s *Suite) Test_CreateGPSUseCase() {
	s.Run("it should be able to create a gps", func() {
		latitude := 490.01
		longitude := -129.92

		input := dto.CreateGPSInputDTO{
			Latitude:  latitude,
			Longitude: longitude,
		}

		output, err := s.useCase.CreateGPSUseCase(input)
		s.NoError(err)

		s.Equal(latitude, output.GPS.Latitude)
		s.Equal(longitude, output.GPS.Longitude)

		repoGPS := s.gpsRepo.Items[0]
		s.Equal(latitude, repoGPS.Latitude)
		s.Equal(longitude, repoGPS.Longitude)
	})
}
