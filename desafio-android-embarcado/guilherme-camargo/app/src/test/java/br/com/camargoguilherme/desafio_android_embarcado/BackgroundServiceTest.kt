package br.com.camargoguilherme.desafio_android_embarcado

import android.app.Notification
import android.app.NotificationManager
import android.content.Context
import android.content.Intent
import android.os.Build
import android.os.Handler
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class BackgroundServiceTest {

    private lateinit var backgroundService: BackgroundService
    private lateinit var context: Context
    private lateinit var handler: Handler

    @Before
    fun setUp() {
        backgroundService = spyk(BackgroundService(), recordPrivateCalls = true)
        context = ApplicationProvider.getApplicationContext()
        handler = mockk(relaxed = true)
    }

    @Test
    fun `onCreate should initialize services`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        backgroundService.onCreate()

        verify { anyConstructed<LogWriter>().writeLog(BackgroundService.TAG, "Service criado") }
    }

    @Test
    fun `onStartCommand should write log when started`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        backgroundService.onStartCommand(Intent(), 0, 1)

        verify { anyConstructed<LogWriter>().writeLog(BackgroundService.TAG, "Serviço iniciado") }
    }

    @Test
    fun `onDestroy should stop services`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        backgroundService.onDestroy()

        verify { anyConstructed<LogWriter>().writeLog(BackgroundService.TAG, "Serviço destruído") }
    }
}
