package com.example.androidembarcado.service

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import com.example.androidembarcado.model.GyroscopeData
import com.example.androidembarcado.repository.TelemetryRepository
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class SensorService(private val context: Context) : SensorEventListener {

    private val sensorManager: SensorManager =
        context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
    private val repository: TelemetryRepository = TelemetryRepository(context)

    fun startGyroscopeTracking(deviceId: String) {
        val gyroscope = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
        sensorManager.registerListener(this, gyroscope, SensorManager.SENSOR_DELAY_NORMAL)
    }

    fun stopGyroscopeTracking() {
        sensorManager.unregisterListener(this)
    }

    override fun onSensorChanged(event: SensorEvent) {
        val gyroscopeData = GyroscopeData(
            x = event.values[0],
            y = event.values[1],
            z = event.values[2],
            timestamp = System.currentTimeMillis(),
            deviceId = "your_device_id" // Substitua pelo ID do dispositivo.
        )

        CoroutineScope(Dispatchers.IO).launch {
            repository.insertGyroscopeData(gyroscopeData)
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {}
}
