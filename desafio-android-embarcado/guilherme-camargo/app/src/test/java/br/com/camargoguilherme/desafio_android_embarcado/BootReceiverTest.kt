package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.content.Intent
import android.os.Build
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class BootReceiverTest {

    private lateinit var bootReceiver: BootReceiver
    private lateinit var context: Context
    private lateinit var intent: Intent

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        bootReceiver = BootReceiver()
        intent = mockk()
        every { intent.action } returns Intent.ACTION_BOOT_COMPLETED
    }

    @Test
    fun `onReceive should start the service on boot completed`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs
        every { context.startService(any()) } returns null
        every { context.startForegroundService(any()) } returns null

        bootReceiver.onReceive(context, intent)

        verify { anyConstructed<LogWriter>().writeLog(BootReceiver.TAG, "Aplicativo atualizado ou instalado") }
        verify { context.startService(any()) }
    }
}
