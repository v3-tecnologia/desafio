package api

import (
	"net/http"
	"time"
	"v3/internal/api/handlers"
	"v3/pkg/httpcore"
	"v3/pkg/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitService() http.Handler {
	router := chi.NewRouter()
	util.InitLogger()

	router.Use(cors.New(httpcore.DefaultCorsOptions).Handler)

	router.Use(middleware.Timeout(200 * time.Second))

	router.Use(middleware.Recoverer)

	router.Use(httpcore.LoggerMiddleware)

	controller := handlers.NewApiController()
	applyRoutes(router, controller)

	return router
}
