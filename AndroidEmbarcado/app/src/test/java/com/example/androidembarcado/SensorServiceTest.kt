package com.example.androidembarcado

import org.junit.Assert.*
import org.junit.Before
import org.junit.Test
import org.mockito.Mock
import org.mockito.Mockito.*
import org.mockito.MockitoAnnotations

class SensorServiceTest {

    @Mock
    private lateinit var sensorService: SensorService

    @Before
    fun setUp() {
        MockitoAnnotations.openMocks(this)
        sensorService = mock(SensorService::class.java)
    }

    @Test
    fun testStartSensor() {
        `when`(sensorService.startSensor()).thenReturn(true)
        val result = sensorService.startSensor()
        assertTrue(result)
        verify(sensorService).startSensor()
    }

    @Test
    fun testStopSensor() {
        `when`(sensorService.stopSensor()).thenReturn(true)
        val result = sensorService.stopSensor()
        assertTrue(result)
        verify(sensorService).stopSensor()
    }

    @Test
    fun testGetSensorData() {
        val expectedData = "sensor data"
        `when`(sensorService.getSensorData()).thenReturn(expectedData)
        val result = sensorService.getSensorData()
        assertEquals(expectedData, result)
        verify(sensorService).getSensorData()
    }
}