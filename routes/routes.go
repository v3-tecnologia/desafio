package routes

import (
	"github/desafio/handlers"
	"github/desafio/repository"
	"github/desafio/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Função que inicia as rotas da API
func InitializeRoutes(router *mux.Router) {

	repository, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err.Error())
	}
	handle := handlers.NewHandle(service.NewService(repository))
	health(router)
	router.HandleFunc("/telemetry/gyroscope", handle.GyroscopeData).Methods("POST")
	router.HandleFunc("/telemetry/gps", handle.GPSData).Methods("POST")
	router.HandleFunc("/telemetry/photo", handle.PhotoData).Methods("POST")
	mux.CORSMethodMiddleware(router)
}

//Função para verificar se há a conexão com o servidor
func health(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(":)"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
