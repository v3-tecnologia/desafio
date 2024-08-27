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

// Função que define as rotas utilizando o gorilla/mux e cria as estruturas requestHandle, service e repository que serão usados pelo servidor
func InitiateRouter() *mux.Router {
	r := mux.NewRouter()

	repository, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	reqHandle := handlers.NewRequestHandle(service.NewService(repository))

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
		log.Fatal(err.Error())
	}

	return r
}
