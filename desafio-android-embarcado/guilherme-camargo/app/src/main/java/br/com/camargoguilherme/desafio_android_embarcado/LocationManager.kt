package br.com.camargoguilherme.desafio_android_embarcado

import android.annotation.SuppressLint
import android.content.Context
import android.location.Location
import android.util.Log
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationCallback
import com.google.android.gms.location.LocationRequest
import com.google.android.gms.location.LocationResult
import com.google.android.gms.location.LocationServices

class LocationManager(context: Context) {

    private var fusedLocationClient: FusedLocationProviderClient = LocationServices.getFusedLocationProviderClient(context)
    private var locationCallback: LocationCallback? = null

    var latitude: Double? = null
        private set
    var longitude: Double? = null
        private set

    @SuppressLint("MissingPermission")
    fun startLocationUpdates() {
        Log.i(BackgroundService.TAG, "startLocationUpdates")
        val locationRequest = LocationRequest.create().apply {
            interval = 10000 // Intervalo para atualizações, em milissegundos
            fastestInterval = 5000 // Intervalo mais rápido para atualizações
            priority = LocationRequest.PRIORITY_HIGH_ACCURACY
        }

        locationCallback = object : LocationCallback() {
            override fun onLocationResult(p0: LocationResult) {
                p0 ?: return
                for (location in p0.locations) {
                    latitude = location.latitude
                    longitude = location.longitude
                }
            }
        }

        fusedLocationClient.requestLocationUpdates(locationRequest, locationCallback!!, null)
    }

    fun stopLocationUpdates() {
        locationCallback?.let {
            fusedLocationClient.removeLocationUpdates(it)
        }
    }
}
