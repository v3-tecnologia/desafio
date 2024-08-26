package main

import (
	"desafio-backend/web/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("starting api desafio-backend")

	router := api.NewAPI()

	api.Start(router)
	logrus.Info("shutting down desafio-backend")
}
