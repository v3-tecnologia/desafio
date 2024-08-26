package com.entregas.android_pleno_teste_v3.backgroundserviceteste

import android.hardware.Sensor
import android.hardware.SensorManager
import android.location.Location
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity
import com.entregas.android_pleno_teste_v3.data.LocationEntity
import com.entregas.android_pleno_teste_v3.data.database.AppDatabase
import com.entregas.android_pleno_teste_v3.services.BackgroundService
import io.mockk.*
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Test

class BackgroundServiceTest {
//caso de teste ainda não está completo mas ja temos uma ideia de como vai ficar

    private lateinit var service: BackgroundService
    private lateinit var mockDatabase: AppDatabase

    @Before
    fun setUp() {
        mockDatabase = mockk()
        service = spyk(BackgroundService())
        service.dataBase = mockDatabase
        every { mockDatabase.gyroscopeDao() } returns mockk()
        every { mockDatabase.locationDao() } returns mockk()
    }

    @Test
    fun `test onLocationChanged stores location data`() = runBlocking {
        val location = mockk<Location>(relaxed = true)
        every { location.latitude } returns 10.0
        every { location.longitude } returns 20.0
        service.onLocationChanged(location)
        //  verify { mockDatabase.locationDao().(LocationEntity(0, 10.0, 20.0, any())) }//
    }

//    @Test
    //fun `test captureSensorData stores gyroscope data`() = runBlocking {
    //     val sensorManager = mockk<SensorManager>()
    //  val sensor = mockk<Sensor>()
    // val gyroscopeDao = mockk()
    //   val timestamp = System.currentTimeMillis()
    //  every { sensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE) } returns sensor
    // every { mockDatabase.gyroscopeDao() } returns gyroscopeDao
    //  service.initializeSensorManager()
    //  service.captureSensorData()
    //  verify { gyroscopeDao.insert(GyroscopeEntity(1, 0f, 0f, 0f, timestamp)) }
    // }
}
