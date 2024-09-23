package routes

import (
    "BackendGo/controllers"
    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/telemetry/gyroscope", controllers.HandleGyroscopeData).Methods("POST")
    router.HandleFunc("/telemetry/gps", controllers.HandleGpsData).Methods("POST")
    router.HandleFunc("/telemetry/photo", controllers.HandlePhotoData).Methods("POST")
    return router
}
