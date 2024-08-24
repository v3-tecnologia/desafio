package com.entregas.android_pleno_teste_v3.data

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "location_data")
data class LocationEntity(
    @PrimaryKey(autoGenerate = true) val id: Int = 0,
    val latitude: Double,
    val longitude: Double,
    val timestamp: Long
)
