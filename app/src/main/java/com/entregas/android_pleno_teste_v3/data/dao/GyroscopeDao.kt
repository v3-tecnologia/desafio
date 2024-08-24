package com.entregas.android_pleno_teste_v3.data.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import androidx.room.Upsert
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity

@Dao
interface GyroscopeDao {
    @Upsert
    suspend fun insert(gyroscope: GyroscopeEntity)

    @Query("SELECT * FROM gyroscope_data")
    suspend fun getAllGyroscopes(): List<GyroscopeEntity>
}
