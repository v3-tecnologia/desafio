package test

import (
	"v3/internal/api/handlers"
	"v3/pkg/httpcore"

	"github.com/go-chi/chi/v5"
)

func setupRouter(endpoint string) *chi.Mux {
	ctl := handlers.NewApiController()
	r := chi.NewRouter()
	r.Post("/telemetry"+endpoint, httpcore.Handle(ctl.CreateGPS))

	return r
}
