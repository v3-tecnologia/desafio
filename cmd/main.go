package main

import (
	"desafio-backend/internal/gyroscope"
	"desafio-backend/web/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("starting api desafio-backend")

	gyroscopeMain := gyroscope.NewMain()

	router := api.NewAPI(gyroscopeMain)

	api.Start(router)
	logrus.Info("shutting down desafio-backend")

}
