package com.example.v3challenge.model

import androidx.annotation.Keep

@Keep
data class Gyro(
    var timestamp: Long? = null,
    var x: String = "",
    var y: String = "",
    var z: String = ""
)

data class Gps(
    var timestamp: Long? = null,
    var lat: String = "",
    var lon: String = ""
)

data class Photo(
    var timestamp: Long? = null,
    var photo: String = ""
)