package com.example.v3challenge.network

import retrofit2.http.Header
import retrofit2.http.POST

interface LogsInterface {

    @POST("/telemetry/gyroscope")
    suspend fun sendGyro(
        @Header("gyro") gyro: String
    ): String

    @POST("/telemetry/gps")
    suspend fun sendGps(
        @Header("gps") gps: String
    ): String

    @POST("/telemetry/photo")
    suspend fun sendPhoto(
        @Header("photo") photo: String
    ): String
}