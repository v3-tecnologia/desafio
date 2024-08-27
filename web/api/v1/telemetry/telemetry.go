package telemetry

import (
	"desafio-backend/internal/gps"
	"desafio-backend/internal/gyroscope"
	"desafio-backend/internal/photo"
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

func Gps(gpsMain gps.UseCases) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the received body to the Request gyroscope struct
		request, err := gpsMain.ParseGps(r.Body)

		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		createdGyroscope, err := gpsMain.SaveGps(request)

		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.NewResponse(w, http.StatusCreated, createdGyroscope)
	}
}

func Photo(photoMain photo.UseCases) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse the received body to the Request Photo upload

		// max total file size 10Mb
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		file, handler, err := r.FormFile("image")

		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		imageFile, errImage := photoMain.ParseImage(file, handler)
		if errImage != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		request, errPhoto := photoMain.ParsePhoto(r.FormValue("request"), imageFile)
		if errPhoto != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		createdGyroscope, errSave := photoMain.SavePhoto(request)

		if errSave != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.NewResponse(w, http.StatusCreated, createdGyroscope)
	}
}
