package com.entregas.android_pleno_teste_v3.services

import android.Manifest
import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Intent
import android.content.pm.PackageManager
import android.content.pm.ServiceInfo.FOREGROUND_SERVICE_TYPE_LOCATION
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import android.location.Location
import android.location.LocationListener
import android.location.LocationManager
import android.os.Build
import android.os.Handler
import android.os.IBinder
import android.os.Looper
import android.util.Log
import androidx.core.app.ActivityCompat
import androidx.core.app.NotificationCompat
import com.entregas.android_pleno_teste_v3.R
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity
import com.entregas.android_pleno_teste_v3.data.LocationEntity
import com.entregas.android_pleno_teste_v3.data.database.AppDatabase
import com.entregas.android_pleno_teste_v3.domain.GPSRequestDataClass
import com.entregas.android_pleno_teste_v3.domain.GyroscopeRequestDataClass
import com.entregas.android_pleno_teste_v3.repository.GiroscopioRequestRepository
import com.entregas.android_pleno_teste_v3.repository.GpsRequestRepository
import com.entregas.android_pleno_teste_v3.utils.Constants
import com.entregas.android_pleno_teste_v3.utils.Constants.Companion.CAPTURE_INTERVAL
import com.entregas.android_pleno_teste_v3.utils.GetMacAddress
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class BackgroundService : Service(), LocationListener {

    private lateinit var locationManager: LocationManager
    private lateinit var sensorManager: SensorManager
    private var gyroscopeSensor: Sensor? = null
    private val gpsRequestRepository = GpsRequestRepository()
    private val giroscopioRequestRepository = GiroscopioRequestRepository()

    lateinit var dataBase: AppDatabase

    private lateinit var coroutineScope: CoroutineScope

    private val handler = Handler(Looper.getMainLooper())

    private var lastGyroscopeValues: FloatArray? = null

    private val gyroscopeEventListener = object : SensorEventListener {
        override fun onSensorChanged(event: SensorEvent?) {
            event?.let {
                lastGyroscopeValues = it.values.copyOf()
            }
        }

        override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {
            // Implement this method if needed
        }
    }

    override fun onCreate() {
        super.onCreate()
        dataBase = AppDatabase.getDatabase(applicationContext)
        coroutineScope = CoroutineScope(Dispatchers.IO)
        initializeNotification()
        initializeLocationManager()
        initializeSensorManager()

        startLocationUpdates()
        registerGyroscopeSensor()

        handler.post(runnableTask)
    }

    private fun initializeNotification() {
        createNotificationChannel()
        val notification: Notification = NotificationCompat.Builder(this, Constants.CHANNEL_ID)
            .setContentTitle("Background Service")
            .setContentText("Service is running...")
            .setSmallIcon(R.drawable.ic_launcher_foreground)
            .build()

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
            startForeground(
                Constants.NOTIFICATION_ID,
                notification,
                FOREGROUND_SERVICE_TYPE_LOCATION
            )
        } else {
            startForeground(Constants.NOTIFICATION_ID, notification)
        }
    }

    private fun initializeLocationManager() {
        locationManager = getSystemService(LOCATION_SERVICE) as LocationManager
    }

    private fun initializeSensorManager() {
        sensorManager = getSystemService(SENSOR_SERVICE) as SensorManager
        gyroscopeSensor = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
    }

    private fun startLocationUpdates() {
        if (ActivityCompat.checkSelfPermission(
                this,
                Manifest.permission.ACCESS_FINE_LOCATION
            ) == PackageManager.PERMISSION_GRANTED || ActivityCompat.checkSelfPermission(
                this,
                Manifest.permission.ACCESS_COARSE_LOCATION
            ) == PackageManager.PERMISSION_GRANTED
        ) {
            locationManager.requestLocationUpdates(
                LocationManager.GPS_PROVIDER,
                CAPTURE_INTERVAL,
                0f,
                this
            )
        }
    }

    private fun registerGyroscopeSensor() {
        gyroscopeSensor?.let {
            sensorManager.registerListener(
                gyroscopeEventListener,
                it,
                SensorManager.SENSOR_DELAY_NORMAL
            )
        }
    }

    private fun sendGyroscopeData(x: Float, y: Float, z: Float, timestamp: Long) {
        coroutineScope.launch {
            dataBase.gyroscopeDao().insert(
                GyroscopeEntity(
                    1,
                    x,
                    y,
                    z,
                    timestamp
                )
            )
            giroscopioRequestRepository.sendGyroscopeRequest(
                GyroscopeRequestDataClass(
                    x,
                    y,
                    z,
                    GetMacAddress(applicationContext).getUniqueDeviceId()
                )
            )
        }
    }

    private fun sendLocationData(latitude: Double, longitude: Double, timestamp: Long) {
        val localEntity = LocationEntity(
            0,
            latitude,
            longitude,
            timestamp
        )
        coroutineScope.launch {
            dataBase.locationDao().insert(
                localEntity
            )
            val gpsRequestDataClass = GPSRequestDataClass(
                latitude,
                longitude,
                GetMacAddress(applicationContext).getUniqueDeviceId()
            )
            gpsRequestRepository.sendGpsRequest(
                gpsRequestDataClass
            )
        }
    }

    private val runnableTask = object : Runnable {
        override fun run() {
            captureSensorData()
            handler.postDelayed(this, 10000) // 10 segundos em milissegundos
        }
    }

    private fun captureSensorData() {
        lastGyroscopeValues?.let { values ->
            val (x, y, z) = values
            val timestamp = System.currentTimeMillis()
            sendGyroscopeData(x, y, z, timestamp)

            // Adicionando logs para ver os valores no Logcat
            Log.d("GyroscopeData", "X: $x, Y: $y, Z: $z, Timestamp: $timestamp")
        }
    }

    override fun onLocationChanged(location: Location) {
        val latitude = location.latitude
        val longitude = location.longitude
        val timestamp = System.currentTimeMillis()
        sendLocationData(latitude, longitude, timestamp)
    }

    override fun onBind(intent: Intent?): IBinder? {
        return null
    }

    private fun createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val name = "Background Service"
            val descriptionText = "Service running in the background"
            val importance = NotificationManager.IMPORTANCE_DEFAULT
            val channel = NotificationChannel(Constants.CHANNEL_ID, name, importance).apply {
                description = descriptionText
            }
            val notificationManager: NotificationManager =
                getSystemService(NotificationManager::class.java)
            notificationManager.createNotificationChannel(channel)
        }
    }

    override fun onDestroy() {
        super.onDestroy()
        handler.removeCallbacks(runnableTask)
        locationManager.removeUpdates(this)
        sensorManager.unregisterListener(gyroscopeEventListener) // Unregister the listener when the service is destroyed
    }
}
