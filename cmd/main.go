package main

import (
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	_ "github.com/HaroldoFV/desafio/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"path/filepath"
)

// @title telemetry Service API
// @version 1.0
// @description This is a telemetry microservice API.
// @host localhost:8080
// @BasePath /api/v1/telemetry/
func main() {
	dir, _ := os.Getwd()
	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Erro ao carregar configurações:", err)
			panic(err)
		}
	}

	webServer, err := InitializeApplication(config)
	if err != nil {
		panic(err)
	}

	webServer.WebServer.AddHandler(http.MethodPost, "/gyroscope", webServer.GyroscopeHandler.Create)
	webServer.WebServer.AddHandler(http.MethodPost, "/gps", webServer.GPSHandler.Create)
	webServer.WebServer.AddHandler(http.MethodPost, "/photo", webServer.PhotoHandler.Create)
	webServer.WebServer.AddHandler(http.MethodGet, "/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.WebServerPort+"/docs/doc.json"),
	))

	fmt.Println("Starting web server on port", config.WebServerPort)
	go func() {
		err = webServer.WebServer.Start()
		if err != nil {
			panic(err)
		}
	}()
	select {}
}
