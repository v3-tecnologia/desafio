package com.entregas.android_pleno_teste_v3.data.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import com.entregas.android_pleno_teste_v3.data.LocationEntity

@Dao
interface LocationDao {
    @Insert
    suspend fun insert(location: LocationEntity)

    @Query("SELECT * FROM location_data")
    suspend fun getAllLocations(): List<LocationEntity>
}
