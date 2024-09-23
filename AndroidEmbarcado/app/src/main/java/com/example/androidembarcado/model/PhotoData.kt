package com.example.androidembarcado.model

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "photo_data")
data class PhotoData(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val photoPath: String,
    val timestamp: Long,
    val deviceId: String
)
