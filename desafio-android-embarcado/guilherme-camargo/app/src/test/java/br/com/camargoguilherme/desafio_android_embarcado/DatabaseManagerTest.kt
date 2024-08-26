package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.database.sqlite.SQLiteDatabase
import androidx.test.core.app.ApplicationProvider
import io.mockk.*
import org.junit.Before
import org.junit.Test

class DatabaseManagerTest {

    private lateinit var databaseManager: DatabaseManager
    private lateinit var context: Context

    @Before
    fun setUp() {
        context = ApplicationProvider.getApplicationContext()
        databaseManager = spyk(DatabaseManager(context))
    }

    @Test
    fun `onCreate should create tables successfully`() {
        mockkConstructor(LogWriter::class)
        val db = mockk<SQLiteDatabase>(relaxed = true)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        databaseManager.onCreate(db)

        verify { db.execSQL(any()) }
        verify { anyConstructed<LogWriter>().writeLog(DatabaseManager.TAG, "Tabelas criadas com sucesso.") }
    }

    @Test
    fun `insertGyroscopeData should log success when inserting data`() {
        mockkConstructor(LogWriter::class)
        every { anyConstructed<LogWriter>().writeLog(any(), any()) } just Runs

        databaseManager.insertGyroscopeData(1f, 2f, 3f, "deviceId")

        verify { anyConstructed<LogWriter>().writeLog(DatabaseManager.TAG, "Dados do giroscópio inseridos com sucesso: ID=1") }
    }
}
