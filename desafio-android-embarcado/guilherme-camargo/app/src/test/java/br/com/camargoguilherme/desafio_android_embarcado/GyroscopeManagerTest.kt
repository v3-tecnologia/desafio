package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorManager
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class GyroscopeManagerTest {

    private lateinit var gyroscopeManager: GyroscopeManager
    private lateinit var context: Context

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        gyroscopeManager = spyk(GyroscopeManager(context))
    }

    @Test
    fun `onSensorChanged should update values when sensor changes`() {
        val event = mockk<SensorEvent>(relaxed = true)
        //event.sensor = mockk(relaxed = true)
        //event.sensor.type = Sensor.TYPE_GYROSCOPE
        every { event.values[0] } returns 1f
        every { event.values[1] } returns 2f
        every { event.values[2] } returns 3f

        gyroscopeManager.onSensorChanged(event)

        assert(gyroscopeManager.xValue == 1f)
        assert(gyroscopeManager.yValue == 2f)
        assert(gyroscopeManager.zValue == 3f)
    }
}
