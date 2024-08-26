package com.entregas.android_pleno_teste_v3.domain.data

data class GpsRequest(
    val latitude: Double,
    val longitude: Double,
    val altitude: Double,
    val macAdd: String
)