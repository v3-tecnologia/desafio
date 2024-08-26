package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.content.Intent
import android.os.Build
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class AppUpdateReceiverTest {

    private lateinit var appUpdateReceiver: AppUpdateReceiver
    private lateinit var context: Context
    private lateinit var intent: Intent
    private lateinit var logWriter: LogWriter

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        appUpdateReceiver = AppUpdateReceiver()
        intent = mockk()
        logWriter = mockk(relaxed = true)
        every { intent.action } returns Intent.ACTION_MY_PACKAGE_REPLACED
    }

    @Test
    fun `onReceive should start the service when action is MY_PACKAGE_REPLACED`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs
        every { context.startService(any()) } returns null
        every { context.startForegroundService(any()) } returns null

        appUpdateReceiver.onReceive(context, intent)

        verify { anyConstructed<LogWriter>().writeLog(AppUpdateReceiver.TAG, "Aplicativo atualizado ou instalado") }
        verify { context.startService(any()) }
    }
}
