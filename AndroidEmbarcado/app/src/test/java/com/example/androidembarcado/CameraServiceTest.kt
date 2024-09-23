package com.example.androidembarcado

import org.junit.Before
import org.junit.Test
import org.mockito.Mock
import org.mockito.Mockito.*
import org.mockito.MockitoAnnotations

class CameraServiceTest {

    @Mock
    private lateinit var cameraService: CameraService

    @Before
    fun setUp() {
        MockitoAnnotations.openMocks(this)
        cameraService = mock(CameraService::class.java)
    }

    @Test
    fun testStartCamera() {
        // Arrange
        doNothing().`when`(cameraService).startCamera()

        // Act
        cameraService.startCamera()

        // Assert
        verify(cameraService, times(1)).startCamera()
    }

    @Test
    fun testStopCamera() {
        // Arrange
        doNothing().`when`(cameraService).stopCamera()

        // Act
        cameraService.stopCamera()

        // Assert
        verify(cameraService, times(1)).stopCamera()
    }
}