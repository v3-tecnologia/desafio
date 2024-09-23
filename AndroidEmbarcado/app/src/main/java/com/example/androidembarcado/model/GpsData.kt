package com.example.androidembarcado.model

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "gps_data")
data class GpsData(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val latitude: Double,
    val longitude: Double,
    val timestamp: Long,
    val deviceId: String
)
