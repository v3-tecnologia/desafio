package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import android.util.Log

class GyroscopeManager(context: Context) : SensorEventListener {

    private val sensorManager: SensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
    private val gyroscope: Sensor? = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
    var xValue: Float = 0f
        private set
    var yValue: Float = 0f
        private set
    var zValue: Float = 0f
        private set

    init {
        // Registra o listener para o giroscópio
        gyroscope?.also { gyro ->
            sensorManager.registerListener(this, gyro, SensorManager.SENSOR_DELAY_NORMAL)
        }
    }

    override fun onSensorChanged(event: SensorEvent?) {
        event?.let {
            Log.i(BackgroundService.TAG, it.sensor.type.toString())
            if (it.sensor.type == Sensor.TYPE_GYROSCOPE) {
                xValue = it.values[0]
                yValue = it.values[1]
                zValue = it.values[2]
            }
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {
        // Não utilizado, mas necessário para a interface SensorEventListener
    }

    fun unregisterListener() {
        // Remove o listener do giroscópio
        sensorManager.unregisterListener(this)
    }
}
