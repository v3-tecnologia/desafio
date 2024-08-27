package main

import (
	"desafio-backend/internal/configuration"
	"desafio-backend/internal/device"
	"desafio-backend/internal/gps"
	"desafio-backend/internal/gyroscope"
	"desafio-backend/internal/photo"
	"desafio-backend/pkg/logger"
	"desafio-backend/web/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("starting api desafio-backend")

	// Crete Postgres Connection
	db, err := configuration.GetDBConnection()
	if err != nil {
		logger.Fatal("Erro ao conectar ao banco de dados", err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	deviceMain := device.NewMain(db)
	gyroscopeMain := gyroscope.NewMain(db, deviceMain)
	gpsMain := gps.NewMain(db, deviceMain)
	photoMain := photo.NewMain(db, deviceMain)

	router := api.NewAPI(gyroscopeMain, gpsMain, photoMain)

	api.Start(router)
	logrus.Info("shutting down desafio-backend")

}
