package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.location.Location
import androidx.test.core.app.ApplicationProvider
import com.google.android.gms.location.LocationAvailability
import com.google.android.gms.location.LocationResult
import io.mockk.*
import org.junit.Before
import org.junit.Test

class LocationManagerTest {

    private lateinit var locationManager: LocationManager
    private lateinit var context: Context

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        locationManager = spyk(LocationManager(context, 1000L, 500f))
    }

    @Test
    fun `onLocationResult should update latitude and longitude`() {
        val location = mockk<Location>(relaxed = true)
        every { location.latitude } returns 1.0
        every { location.longitude } returns 2.0
        val locationResult = LocationResult.create(listOf(location))

        locationManager.onLocationResult(locationResult)

        assert(locationManager.latitude == 1.0)
        assert(locationManager.longitude == 2.0)
    }

    @Test
    fun `onLocationAvailability should log availability`() {
        val locationAvailability = mockk<LocationAvailability>(relaxed = true)
        every { locationAvailability.isLocationAvailable } returns true
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        locationManager.onLocationAvailability(locationAvailability)

        verify { anyConstructed<LogWriter>().writeLog(LocationManager.TAG, "Disponibilidade de localização: true") }
    }
}
