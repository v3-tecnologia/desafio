package multipart

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/shared"
)

type Handler interface {
	Read(request shared.Request) (*domain.File, error)
}

type Adapter struct {
	handler Handler
}

type AdapterInterface interface {
	Read(request *shared.Request) (*domain.File, error)
}

func NewMultipartAdapter(handler Handler) *Adapter {
	return &Adapter{
		handler: handler,
	}
}

func (m *Adapter) Reader(request shared.Request) (*domain.File, error) {
	return m.handler.Read(request)
}
