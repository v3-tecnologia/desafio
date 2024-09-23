package com.example.androidembarcado

import org.junit.Assert.*
import org.junit.Before
import org.junit.Test
import org.mockito.Mockito.*

class GpsServiceTest {

    private lateinit var gpsService: GpsService
    private lateinit var locationProvider: LocationProvider

    @Before
    fun setUp() {
        locationProvider = mock(LocationProvider::class.java)
        gpsService = GpsService(locationProvider)
    }

    @Test
    fun testGetCurrentLocation() {
        val expectedLocation = Location(12.34, 56.78)
        `when`(locationProvider.getLocation()).thenReturn(expectedLocation)

        val actualLocation = gpsService.getCurrentLocation()

        assertEquals(expectedLocation, actualLocation)
    }

    @Test
    fun testIsGpsEnabled() {
        `when`(locationProvider.isGpsEnabled()).thenReturn(true)

        val isGpsEnabled = gpsService.isGpsEnabled()

        assertTrue(isGpsEnabled)
    }
}