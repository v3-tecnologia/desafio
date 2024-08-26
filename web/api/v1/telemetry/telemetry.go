package telemetry

import (
	"desafio-backend/web/api/util"
	"net/http"
)

func Gyroscope() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.NewResponse(w, http.StatusOK, "Dados de giroscópio recebidos com sucesso")
	}
}

func Gps() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.NewResponse(w, http.StatusOK, "Dados de gps recebidos com sucesso")
	}
}

func Photo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.NewResponse(w, http.StatusOK, "Dados de foto recebidos com sucesso")
	}
}
