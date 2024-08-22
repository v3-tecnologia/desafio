package br.com.camargoguilherme.desafio_android_embarcado

import android.Manifest
import android.annotation.SuppressLint
import android.content.Context
import android.content.pm.PackageManager
import android.os.Looper
import androidx.core.app.ActivityCompat
import androidx.core.content.ContextCompat
import com.google.android.gms.location.*

// Exemplo from https://tomas-repcik.medium.com/locationrequest-create-got-deprecated-how-to-fix-it-e4f814138764
class LocationManager(private val context: Context,
                      private var timeInterval: Long,
                      private var minimalDistance: Float) : LocationCallback() {
    companion object {
        const val TAG = "LocationManager"
    }

    private var request: LocationRequest = createRequest()
    private var locationClient: FusedLocationProviderClient = LocationServices.getFusedLocationProviderClient(context)
    private val logWriter = LogWriter(context)

    var latitude: Double? = null
        private set
    var longitude: Double? = null
        private set

    private fun createRequest(): LocationRequest =
        // New builder
        LocationRequest.Builder(Priority.PRIORITY_HIGH_ACCURACY, timeInterval).apply {
            setMinUpdateDistanceMeters(minimalDistance)
            setGranularity(Granularity.GRANULARITY_PERMISSION_LEVEL)
            setWaitForAccurateLocation(true)
        }.build()

    fun changeRequest(timeInterval: Long, minimalDistance: Float) {
        this.timeInterval = timeInterval
        this.minimalDistance = minimalDistance
        createRequest()
        stopLocationTracking()
        startLocationTracking()
    }

    fun checkPermissions(): Boolean {
        val fineLocationPermission = ContextCompat.checkSelfPermission(context, Manifest.permission.ACCESS_FINE_LOCATION)
        val coarseLocationPermission = ContextCompat.checkSelfPermission(context, Manifest.permission.ACCESS_COARSE_LOCATION)
        return fineLocationPermission == PackageManager.PERMISSION_GRANTED &&
                coarseLocationPermission == PackageManager.PERMISSION_GRANTED
    }

    @SuppressLint("MissingPermission")
    fun startLocationTracking() {
        try {
            // Verificação de permissões
            if (!checkPermissions()) {
                logWriter.writeLog(TAG, "Permissões de localização não concedidas")
                return
            }
            locationClient.requestLocationUpdates(request, this, Looper.getMainLooper())
            logWriter.writeLog(TAG, "Atualizações de localização iniciadas com sucesso")
        } catch (e: SecurityException) {
            logWriter.writeLog(TAG, "Permissão de localização não concedida: ${e.message}")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao iniciar atualizações de localização: ${e.message}")
        }
    }

    fun stopLocationTracking() {
        try {
            locationClient.flushLocations()
            locationClient.removeLocationUpdates(this)
            logWriter.writeLog(TAG, "Atualizações de localização paradas com sucesso")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao parar atualizações de localização: ${e.message}")
        }
    }

    override fun onLocationResult(location: LocationResult) {
        if (location == null) {
            logWriter.writeLog(TAG, "LocationResult é nulo")
            return
        }
        for (location in location.locations) {
            latitude = location.latitude
            longitude = location.longitude
            logWriter.writeLog(TAG, "Localização atualizada: ${location.latitude}, ${location.longitude}")
        }
    }

    override fun onLocationAvailability(availability: LocationAvailability) {
        logWriter.writeLog(TAG, "Disponibilidade de localização: ${availability.isLocationAvailable}")
    }
}
