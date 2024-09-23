package com.example.androidembarcado.worker

import android.content.Context
import androidx.work.CoroutineWorker
import androidx.work.WorkerParameters
import com.example.androidembarcado.model.GpsData
import com.example.androidembarcado.model.GyroscopeData
import com.example.androidembarcado.network.NetworkModule
import com.example.androidembarcado.repository.TelemetryRepository
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

class TelemetryWorker(private val context: Context, params: WorkerParameters) :
    CoroutineWorker(context, params) {

    private val repository = TelemetryRepository(context)
    private val apiService = NetworkModule.apiService

    override suspend fun doWork(): Result {
        return withContext(Dispatchers.IO) {
            try {
                val gyroscopeData = repository.getAllGyroscopeData().value?.last()
                val gpsData = repository.getAllGpsData().value?.last()
                val photoData = repository.getAllPhotoData().value?.last()

                gyroscopeData?.let { apiService.sendGyroscopeData(it) }
                gpsData?.let { apiService.sendGpsData(it) }
                photoData?.let { apiService.sendPhotoData(it) }

                Result.success()
            } catch (e: Exception) {
                Result.failure()
            }
        }
    }
}
