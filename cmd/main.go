package main

import (
	"desafio-backend/internal/gps"
	"desafio-backend/internal/gyroscope"
	"desafio-backend/web/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("starting api desafio-backend")

	gyroscopeMain := gyroscope.NewMain()
	gpsMain := gps.NewMain()

	router := api.NewAPI(gyroscopeMain, gpsMain)

	api.Start(router)
	logrus.Info("shutting down desafio-backend")

}
