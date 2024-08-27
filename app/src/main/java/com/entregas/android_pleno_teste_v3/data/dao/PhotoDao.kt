package com.entregas.android_pleno_teste_v3.data.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import com.entregas.android_pleno_teste_v3.data.PhotoEntity

@Dao
interface PhotoDao {
    @Insert
    suspend fun insert(photoEntity: PhotoEntity)

    @Query("SELECT * FROM photo_data")
    suspend fun getAllPhotos(): List<PhotoEntity>
}