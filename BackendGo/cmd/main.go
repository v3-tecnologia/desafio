package main

import (
    "log"
    "net/http"
    "BackendGo/config"
    "BackendGo/routes"
)

func main() {
    // Carregar configuração
    config.LoadConfig()

    // Definir rotas
    router := routes.SetupRoutes()

    // Iniciar servidor
    log.Println("Starting server on port", config.Config.ServerPort)
    log.Fatal(http.ListenAndServe(":"+config.Config.ServerPort, router))
}
