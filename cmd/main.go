package main

import (
	"database/sql"
	"fmt"
	"github.com/HaroldoFV/desafio/configs"
	_ "github.com/HaroldoFV/desafio/docs"
	"github.com/HaroldoFV/desafio/internal/infra/database"
	"github.com/HaroldoFV/desafio/internal/usecase"
	"github.com/HaroldoFV/desafio/internal/web"
	"github.com/HaroldoFV/desafio/internal/web/webserver"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"path/filepath"
)

// @title telemetry Service API
// @version 1.0
// @description This is a telemetry microservice API.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	dir, _ := os.Getwd()
	fmt.Println("Diretório atual:", dir)

	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Erro ao carregar configurações:", err)
			panic(err)
		}
	}
	fmt.Printf("Configurações carregadas: %+v\n", config)

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	webServer := webserver.NewWebServer(":" + config.WebServerPort)

	gyroscopeRepository := database.NewGyroscopeRepository(db)
	createGyroscopeUseCase := usecase.NewCreateGyroscopeUseCase(gyroscopeRepository)
	webGyroscopeHandler := web.NewGyroscopeHandler(createGyroscopeUseCase, gyroscopeRepository)

	webServer.AddHandler(http.MethodPost, "/gyroscopes", webGyroscopeHandler.Create)

	webServer.AddHandler(http.MethodGet, "/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.WebServerPort+"/docs/doc.json"),
	))

	fmt.Println("Starting web server on port", config.WebServerPort)
	go func() {
		err = webServer.Start()
		if err != nil {
			panic(err)
		}
	}()
	select {}
}
