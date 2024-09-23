package com.example.androidembarcado.model

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "gyroscope_data")
data class GyroscopeData(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val x: Float,
    val y: Float,
    val z: Float,
    val timestamp: Long,
    val deviceId: String
)
