package usecase

import "github.com/charmingruby/g3/internal/telemetry/domain/dto"

func (s *Suite) Test_CreateGyroscopeUseCase() {
	s.Run("it should be able to create a gyroscope", func() {
		x := 122.9
		y := 12.999
		z := -145.0

		input := dto.CreateGyroscopeInputDTO{
			XPosition: x,
			YPosition: y,
			ZPosition: z,
		}

		output, err := s.useCase.CreateGyroscopeUseCase(input)
		s.NoError(err)

		s.Equal(x, output.Gyroscope.XPosition)
		s.Equal(y, output.Gyroscope.YPosition)
		s.Equal(z, output.Gyroscope.ZPosition)

		repoGyroscope := s.gyroscopeRepo.Items[0]
		s.Equal(x, repoGyroscope.XPosition)
		s.Equal(y, repoGyroscope.YPosition)
		s.Equal(z, repoGyroscope.ZPosition)
	})
}
