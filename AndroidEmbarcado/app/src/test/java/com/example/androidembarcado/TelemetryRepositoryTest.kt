package com.example.androidembarcado

import org.junit.Assert.*
import org.junit.Before
import org.junit.Test
import org.mockito.Mock
import org.mockito.Mockito.*
import org.mockito.MockitoAnnotations

class TelemetryRepositoryTest {

    @Mock
    private lateinit var mockDataSource: TelemetryDataSource

    private lateinit var telemetryRepository: TelemetryRepository

    @Before
    fun setUp() {
        MockitoAnnotations.openMocks(this)
        telemetryRepository = TelemetryRepository(mockDataSource)
    }

    @Test
    fun testGetTelemetryData() {
        // Arrange
        val expectedData = TelemetryData("example data")
        `when`(mockDataSource.getTelemetryData()).thenReturn(expectedData)

        // Act
        val actualData = telemetryRepository.getTelemetryData()

        // Assert
        assertEquals(expectedData, actualData)
        verify(mockDataSource).getTelemetryData()
    }

    @Test
    fun testSaveTelemetryData() {
        // Arrange
        val dataToSave = TelemetryData("example data")

        // Act
        telemetryRepository.saveTelemetryData(dataToSave)

        // Assert
        verify(mockDataSource).saveTelemetryData(dataToSave)
    }
}