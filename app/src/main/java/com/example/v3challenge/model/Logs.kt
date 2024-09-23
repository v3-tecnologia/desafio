package com.example.v3challenge.model

import android.graphics.Bitmap
import android.os.Parcelable
import androidx.annotation.Keep
import com.example.v3challenge.utils.LogType
import com.squareup.moshi.JsonClass
import kotlinx.parcelize.Parcelize
import kotlinx.parcelize.RawValue

@Keep
@Parcelize
@JsonClass(generateAdapter = true)
data class Gyro(
    var timestamp: Long? = null,
    var x: String = "",
    var y: String = "",
    var z: String = ""
) : Parcelable

@Keep
@Parcelize
@JsonClass(generateAdapter = true)
data class Gps(
    var timestamp: Long? = null,
    var lat: Double? = null,
    var lon: Double? = null
) : Parcelable

@Keep
@Parcelize
@JsonClass(generateAdapter = true)
data class Photo(
    var timestamp: Long? = null,
    var photo: String? = null
) : Parcelable

@Keep
@Parcelize
@JsonClass(generateAdapter = true)
data class GenericLog(
    var type: LogType,
    var log: @RawValue Any
) : Parcelable