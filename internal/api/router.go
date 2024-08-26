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
	"github.com/go-chi/httprate"
)

func InitService() http.Handler {
	router := chi.NewRouter()
	util.InitLogger()

	router.Use(cors.New(httpcore.DefaultCorsOptions).Handler)

	//Timeout longo só pra facilitar debugging
	router.Use(middleware.Timeout(400 * time.Second))

	router.Use(middleware.Recoverer)

	//Permitindo até 60 requests por minuto
	//antes do throttling colocar os próximos requests na fila
	router.Use(middleware.Throttle(60))

	//Cada IP pode fazer 30 requests por minuto
	router.Use(httprate.LimitByIP(30, 1*time.Minute))

	router.Use(httpcore.LoggerMiddleware)

	controller := handlers.NewApiController()
	applyRoutes(router, controller)

	return router
}
