package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager

class GyroscopeManager(context: Context) : SensorEventListener {
    companion object {
        const val TAG = "GyroscopeManager"
    }

    private val sensorManager: SensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
    private val gyroscope: Sensor? = sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE)
    private val logWriter = LogWriter(context)

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
            logWriter.writeLog(TAG,"Giroscópio registrado com sucesso")
        } ?: run {
            logWriter.writeLog(TAG,"Sensor de giroscópio não disponível")
        }
    }

    override fun onSensorChanged(event: SensorEvent?) {
        event?.let {
            if (it.sensor.type == Sensor.TYPE_GYROSCOPE) {
                xValue = it.values[0]
                yValue = it.values[1]
                zValue = it.values[2]
                //logWriter.writeLog(TAG,"Giroscópio - X: $xValue, Y: $yValue, Z: $zValue")
            }
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {
        // Não utilizado, mas necessário para a interface SensorEventListener
    }

    fun unregisterListener() {
        // Remove o listener do giroscópio
        sensorManager.unregisterListener(this)
        logWriter.writeLog(TAG,"Listener do giroscópio removido")
    }
}
