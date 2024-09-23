package com.example.androidembarcado.repository

import android.content.Context
import androidx.lifecycle.LiveData
import com.example.androidembarcado.database.AppDatabase
import com.example.androidembarcado.model.GpsData
import com.example.androidembarcado.model.GyroscopeData
import com.example.androidembarcado.model.PhotoData

class TelemetryRepository(context: Context) {

    private val db = AppDatabase.getInstance(context)

    suspend fun insertGyroscopeData(data: GyroscopeData) {
        db.gyroscopeDao().insert(data)
    }

    suspend fun insertGpsData(data: GpsData) {
        db.gpsDao().insert(data)
    }

    suspend fun insertPhotoData(data: PhotoData) {
        db.photoDao().insert(data)
    }

    fun getAllGyroscopeData(): LiveData<List<GyroscopeData>> {
        return db.gyroscopeDao().getAll()
    }

    fun getAllGpsData(): LiveData<List<GpsData>> {
        return db.gpsDao().getAll()
    }

    fun getAllPhotoData(): LiveData<List<PhotoData>> {
        return db.photoDao().getAll()
    }
}
