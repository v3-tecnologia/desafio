package api

import (
	_ "v3/docs"
	"v3/internal/api/handlers"
	"v3/pkg/httpcore"

	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger"
)

func applyRoutes(r chi.Router, ctl *handlers.ApiController) {
	r.Post("/telemetry/gyroscope", httpcore.Handle(ctl.CreateGyroscope))
	r.Post("/telemetry/gps", httpcore.Handle(ctl.CreateGPS))
	r.Post("/telemetry/photo", httpcore.Handle(ctl.CreatePhoto))

	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
