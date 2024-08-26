package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.hardware.camera2.CameraDevice
import android.hardware.camera2.CameraManager
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class CustomCameraManagerTest {

    private lateinit var customCameraManager: CustomCameraManager
    private lateinit var context: Context

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        customCameraManager = spyk(CustomCameraManager(context))
    }

    @Test
    fun `initializeCameraAndTakePicture should log success on camera open`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        customCameraManager.initializeCameraAndTakePicture()

        verify { anyConstructed<LogWriter>().writeLog(CustomCameraManager.TAG, "Câmera aberta com sucesso") }
    }
}
