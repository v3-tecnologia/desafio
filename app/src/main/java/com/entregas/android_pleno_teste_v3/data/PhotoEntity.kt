package com.entregas.android_pleno_teste_v3.data

import androidx.room.Entity
import androidx.room.PrimaryKey


@Entity(tableName = "photo_data")
data class PhotoEntity (
    @PrimaryKey(autoGenerate = true)
    val id :Int,
    val base64Image: String ,
)