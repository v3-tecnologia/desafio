package com.entregas.android_pleno_teste_v3.services.apiservice

import com.entregas.android_pleno_teste_v3.domain.PhotoRequestDataClass
import com.entregas.android_pleno_teste_v3.domain.GPSRequestDataClass
import com.entregas.android_pleno_teste_v3.domain.GyroscopeRequestDataClass
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiService {
    @POST("/telemetry/gps")
    suspend fun enviarDadosGps(@Body gpsRequest: GPSRequestDataClass): Result<Unit>

    @POST("/telemetry/gyroscope")
    suspend fun enviarDadosGiroscopio(@Body gyroscopeRequest: GyroscopeRequestDataClass): Result<Unit>

    @POST("/telemetry/photo")
    suspend fun enviarFoto(@Body photoRequest: PhotoRequestDataClass): Result<Unit>
}