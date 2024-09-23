package com.example.androidembarcado.service

import android.annotation.SuppressLint
import android.content.Context
import android.location.Location
import android.location.LocationListener
import android.location.LocationManager
import com.example.androidembarcado.model.GpsData
import com.example.androidembarcado.repository.TelemetryRepository
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class GpsService(private val context: Context) : LocationListener {

    private val locationManager: LocationManager =
        context.getSystemService(Context.LOCATION_SERVICE) as LocationManager
    private val repository: TelemetryRepository = TelemetryRepository(context)

    @SuppressLint("MissingPermission")
    fun startGpsTracking(deviceId: String) {
        locationManager.requestLocationUpdates(
            LocationManager.GPS_PROVIDER, 
            10000, 
            0f, 
            this
        )
    }

    fun stopGpsTracking() {
        locationManager.removeUpdates(this)
    }

    override fun onLocationChanged(location: Location) {
        val gpsData = GpsData(
            latitude = location.latitude,
            longitude = location.longitude,
            timestamp = System.currentTimeMillis(),
            deviceId = "your_device_id" // Substitua pelo ID do dispositivo.
        )

        CoroutineScope(Dispatchers.IO).launch {
            repository.insertGpsData(gpsData)
        }
    }
}
