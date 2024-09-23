package com.example.androidembarcado.network

import com.example.androidembarcado.model.GpsData
import com.example.androidembarcado.model.GyroscopeData
import com.example.androidembarcado.model.PhotoData
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiService {

    @POST("/telemetry/gyroscope")
    suspend fun sendGyroscopeData(@Body data: GyroscopeData)

    @POST("/telemetry/gps")
    suspend fun sendGpsData(@Body data: GpsData)

    @POST("/telemetry/photo")
    suspend fun sendPhotoData(@Body data: PhotoData)
}
