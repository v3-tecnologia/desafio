package inmemory_adapter

import "github.com/charmingruby/g3/internal/telemetry/domain/port"

func NewRecognizerInMemoryAdapter() *RecognizerInMemoryAdapter {
	return &RecognizerInMemoryAdapter{
		MockedFaces: []port.DetectedFace{},
	}
}

type RecognizerInMemoryAdapter struct {
	MockedFaces []port.DetectedFace
}

func (r *RecognizerInMemoryAdapter) Recognize(imageURL string) ([]port.DetectedFace, error) {
	return r.MockedFaces, nil
}
