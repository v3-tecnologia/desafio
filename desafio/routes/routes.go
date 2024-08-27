package routes

import (
	"desafio/handlers"
	"desafio/repository"
	"desafio/service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func InitiateRouter() *mux.Router {
	r := mux.NewRouter()

	repository, err := repository.NewRepository()
	if err != nil {
		log.Fatalf(err.Error())
	}

	reqHandle := handlers.NewRequestHandle(service.NewService(repository))

	r.HandleFunc("/", reqHandle.HealthCheck).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/telemetry/gyroscope", reqHandle.GyroscopeHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/telemetry/gps", reqHandle.GpsHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/telemetry/photo", reqHandle.PhotoHandler).Methods(http.MethodPost, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	err = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	return r
}
