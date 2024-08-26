package telemetry

import (
	"desafio-backend/internal/gyroscope"
	"desafio-backend/web/api/util"
	"net/http"
)

func Gyroscope(gyroscopeMain gyroscope.UseCases) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the received body to the Request gyroscope struct
		request, err := gyroscopeMain.ParseGyroscope(r.Body)

		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		createdGyroscope, err := gyroscopeMain.SaveGyroscope(request)

		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.NewResponse(w, http.StatusCreated, createdGyroscope)
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
