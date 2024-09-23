package controllers

import (
    "BackendGo/models"
    "BackendGo/repositories"
    "BackendGo/services"
    "encoding/json"
    "net/http"
)

func HandleGyroscopeData(w http.ResponseWriter, r *http.Request) {
    var data models.GyroscopeData
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := repositories.InsertGyroscopeData(data); err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func HandleGpsData(w http.ResponseWriter, r *http.Request) {
    var data models.GpsData
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := repositories.InsertGpsData(data); err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func HandlePhotoData(w http.ResponseWriter, r *http.Request) {
    var data models.PhotoData
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Verificação com AWS Rekognition
    recognized, err := services.ComparePhotoWithRekognition(data.PhotoPath)
    if err != nil {
        http.Error(w, "AWS Rekognition error", http.StatusInternalServerError)
        return
    }

    if err := repositories.InsertPhotoData(data); err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    response := map[string]bool{"recognized": recognized}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
