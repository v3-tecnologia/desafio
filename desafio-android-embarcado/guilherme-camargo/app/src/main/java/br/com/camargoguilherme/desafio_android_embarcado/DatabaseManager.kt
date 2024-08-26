package br.com.camargoguilherme.desafio_android_embarcado

import android.content.ContentValues
import android.content.Context
import android.database.sqlite.SQLiteDatabase
import android.database.sqlite.SQLiteOpenHelper
import android.util.Log

class DatabaseManager(context: Context) : SQLiteOpenHelper(context, DATABASE_NAME, null, DATABASE_VERSION) {
    companion object {
        const val TAG = "DatabaseManager"
        private const val DATABASE_NAME = "telemetry.db"
        private const val DATABASE_VERSION = 1

        // Nomes das tabelas
        private const val TABLE_GYROSCOPE = "GyroscopeData"
        private const val TABLE_LOCATION = "LocationData"
        private const val TABLE_IMAGE = "ImageData"

        // Coluna comum a todas as tabelas
        private const val COLUMN_ID = "id"
        private const val COLUMN_TIMESTAMP = "timestamp"
        private const val COLUMN_DEVICE_ID = "device_id"

        // Colunas da tabela GyroscopeData
        private const val COLUMN_X_VALUE = "x_value"
        private const val COLUMN_Y_VALUE = "y_value"
        private const val COLUMN_Z_VALUE = "z_value"

        // Colunas da tabela LocationData
        private const val COLUMN_LATITUDE = "latitude"
        private const val COLUMN_LONGITUDE = "longitude"

        // Colunas da tabela ImageData
        private const val COLUMN_IMAGE_BASE64 = "image_base64"
    }

    private val logWriter = LogWriter(context)

    override fun onCreate(db: SQLiteDatabase?) {
        try {
            // Criação da tabela GyroscopeData
            val createGyroscopeTable = """
                CREATE TABLE $TABLE_GYROSCOPE (
                    $COLUMN_ID INTEGER PRIMARY KEY AUTOINCREMENT,
                    $COLUMN_TIMESTAMP INTEGER,
                    $COLUMN_DEVICE_ID STRING,
                    $COLUMN_X_VALUE REAL,
                    $COLUMN_Y_VALUE REAL,
                    $COLUMN_Z_VALUE REAL
                )
            """.trimIndent()

            // Criação da tabela LocationData
            val createLocationTable = """
                CREATE TABLE $TABLE_LOCATION (
                    $COLUMN_ID INTEGER PRIMARY KEY AUTOINCREMENT,
                    $COLUMN_TIMESTAMP INTEGER,
                    $COLUMN_DEVICE_ID STRING,
                    $COLUMN_LATITUDE REAL,
                    $COLUMN_LONGITUDE REAL
                )
            """.trimIndent()

            // Criação da tabela ImageData
            val createImageTable = """
                CREATE TABLE $TABLE_IMAGE (
                    $COLUMN_ID INTEGER PRIMARY KEY AUTOINCREMENT,
                    $COLUMN_TIMESTAMP INTEGER,
                    $COLUMN_DEVICE_ID STRING,
                    $COLUMN_IMAGE_BASE64 TEXT
                )
            """.trimIndent()

            db?.execSQL(createGyroscopeTable)
            db?.execSQL(createLocationTable)
            db?.execSQL(createImageTable)

            logWriter.writeLog(TAG, "Tabelas criadas com sucesso.")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao criar as tabelas: ${e.message}")
        }
    }

    override fun onUpgrade(db: SQLiteDatabase?, oldVersion: Int, newVersion: Int) {
        try {
            db?.execSQL("DROP TABLE IF EXISTS $TABLE_GYROSCOPE")
            db?.execSQL("DROP TABLE IF EXISTS $TABLE_LOCATION")
            db?.execSQL("DROP TABLE IF EXISTS $TABLE_IMAGE")
            onCreate(db)
            logWriter.writeLog(TAG, "Banco de dados atualizado com sucesso.")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao atualizar o banco de dados: ${e.message}")
        }
    }

    // Métodos para inserir dados nas tabelas com tratamento de exceções e logs

    fun insertGyroscopeData(xValue: Float, yValue: Float, zValue: Float, timestamp: Long, deviceId: String): Long {
        return try {
            val db = this.writableDatabase
            val values = ContentValues().apply {
                put(COLUMN_TIMESTAMP, timestamp)
                put(COLUMN_DEVICE_ID, deviceId)
                put(COLUMN_X_VALUE, xValue)
                put(COLUMN_Y_VALUE, yValue)
                put(COLUMN_Z_VALUE, zValue)
            }
            val id = db.insert(TABLE_GYROSCOPE, null, values)
            logWriter.writeLog(TAG, "Dados do giroscópio inseridos com sucesso: ID=$id")
            id
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inserir dados do giroscópio: ${e.message}")
            -1
        }
    }

    fun insertLocationData(latitude: Double, longitude: Double, timestamp: Long, deviceId: String): Long {
        return try {
            val db = this.writableDatabase
            val values = ContentValues().apply {
                put(COLUMN_TIMESTAMP, timestamp)
                put(COLUMN_DEVICE_ID, deviceId)
                put(COLUMN_LATITUDE, latitude)
                put(COLUMN_LONGITUDE, longitude)
            }
            val id = db.insert(TABLE_LOCATION, null, values)
            logWriter.writeLog(TAG, "Dados de localização inseridos com sucesso: ID=$id")
            id
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inserir dados de localização: ${e.message}")
            -1
        }
    }

    fun insertImageData(imageBase64: String, timestamp: Long, deviceId: String): Long {
        return try {
            val db = this.writableDatabase
            val values = ContentValues().apply {
                put(COLUMN_TIMESTAMP, timestamp)
                put(COLUMN_DEVICE_ID, deviceId)
                put(COLUMN_IMAGE_BASE64, imageBase64)
            }
            val id = db.insert(TABLE_IMAGE, null, values)
            logWriter.writeLog(TAG, "Dados da imagem inseridos com sucesso: ID=$id")
            id
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inserir dados da imagem: ${e.message}")
            -1
        }
    }
}
