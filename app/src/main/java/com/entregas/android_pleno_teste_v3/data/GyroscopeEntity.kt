package com.entregas.android_pleno_teste_v3.data

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "gyroscope_data")
data class GyroscopeEntity(
    @PrimaryKey(autoGenerate = true)
    val id: Int = 0,
    val x: Float,
    val y: Float,
    val z: Float,
    val timestamp: Long
)
