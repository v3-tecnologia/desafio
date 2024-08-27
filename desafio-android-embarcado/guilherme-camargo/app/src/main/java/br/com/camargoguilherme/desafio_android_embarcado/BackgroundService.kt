package br.com.camargoguilherme.desafio_android_embarcado

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Context
import android.content.Intent
import android.os.Build
import android.os.Handler
import android.os.IBinder
import android.os.Looper
import android.provider.Settings
import androidx.core.app.NotificationCompat
import kotlinx.coroutines.*
import org.json.JSONObject

class BackgroundService : Service() {

    companion object {
        const val CHANNEL_ID = "BackgroundServiceChannel"
        const val TAG = "BackgroundService"
        const val BASE_URL = "https://your-api-url.com/telemetry"
        const val GYROSCOPE_ENDPOINT = "$BASE_URL/gyroscope"
        const val GPS_ENDPOINT = "$BASE_URL/gps"
        const val PHOTO_ENDPOINT = "$BASE_URL/photo"
    }

    private lateinit var logWriter: LogWriter
    private lateinit var apiClient: ApiClient
    private lateinit var databaseManager: DatabaseManager
    private lateinit var gyroscopeManager: GyroscopeManager
    private lateinit var locationManager: LocationManager
    private lateinit var customCameraManager: CustomCameraManager
    private lateinit var deviceId: String
    private val handler = Handler(Looper.getMainLooper())
    private val interval: Long = 10000 // 10 segundos
    private val coroutineScope = CoroutineScope(Dispatchers.IO + Job())

    private val runnable = object : Runnable {
        override fun run() {
            coroutineScope.launch {
                try {
                    executeRoutine()
                } catch (e: Exception) {
                    logWriter.writeLog(TAG, "Erro na execução da rotina: ${e.message}")
                }
            }
            handler.postDelayed(this, interval)
        }
    }

    override fun onCreate() {
        super.onCreate()

        logWriter = LogWriter(this)
        logWriter.writeLog(TAG, "Service criado")

        // Captura a identificação única do dispositivo
        deviceId = getAndroidId()

        // Inicializa o ApiClient
        try {
            apiClient = ApiClient(this)
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inicializar ApiClient: ${e.message}")
        }

        // Inicializa o DatabaseManager
        try {
            databaseManager = DatabaseManager(this)
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inicializar DatabaseManager: ${e.message}")
        }

        // Inicializa o GyroscopeManager, LocationManager e CustomCameraManager
        try {
            gyroscopeManager = GyroscopeManager(this)
            locationManager = LocationManager(this, 1000, 500f)
            customCameraManager = CustomCameraManager(this)
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inicializar GyroscopeManager, LocationManager ou CustomCameraManager: ${e.message}")
        }

        // Inicia as atualizações de localização
        try {
            locationManager.startLocationTracking()
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao iniciar atualizações de localização: ${e.message}")
        }

        // Inicia a rotina imediatamente
        handler.post(runnable)
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        logWriter.writeLog(TAG, "Serviço iniciado")
        return START_STICKY
    }

    override fun onDestroy() {
        super.onDestroy()
        logWriter.writeLog(TAG, "Serviço destruído")

        // Cancela todas as corrotinas ao destruir o serviço
        coroutineScope.cancel()

        // Para as atualizações de localização e remove as callbacks da rotina
        try {
            locationManager.stopLocationTracking()
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao parar atualizações de localização: ${e.message}")
        }

        handler.removeCallbacks(runnable)
    }

    override fun onBind(intent: Intent?): IBinder? {
        return null
    }

    private suspend fun executeRoutine() = coroutineScope {
        val gyroscopeJob = launch { getDataGyroscope() }
        val locationJob = launch { getDataLocation() }
        val cameraJob = launch { getDataCamera() }

        // Espera todas as tarefas terminarem antes de continuar
        joinAll(gyroscopeJob, locationJob, cameraJob)
    }

    private suspend fun getDataGyroscope() {
        withContext(Dispatchers.IO) {
            try {
                // Acessa os valores do giroscópio
                val x = gyroscopeManager.xValue
                val y = gyroscopeManager.yValue
                val z = gyroscopeManager.zValue

                // Gerar o timestamp
                val timestamp = System.currentTimeMillis()

                // Loga os valores do giroscópio
                val messageGyroscope = "Giroscópio na rotina - X: $x, Y: $y, Z: $z"
                logWriter.writeLog(TAG, messageGyroscope)

                if (x != null && y != null && z != null) {
                    val json = JSONObject().apply {
                        put("timestamp", timestamp)
                        put("x", x)
                        put("y", y)
                        put("z", z)
                    }
                    apiClient.postSendData(GYROSCOPE_ENDPOINT, json.toString())
                    databaseManager.insertGyroscopeData(x, y, z, timestamp, deviceId)
                }
            } catch (e: Exception) {
                logWriter.writeLog(TAG, "Erro ao buscar dados do giroscópio: ${e.message}")
            }
        }
    }

    private suspend fun getDataLocation() {
        withContext(Dispatchers.IO) {
            try {
                // Acessa os valores de localização
                val latitude = locationManager.latitude
                val longitude = locationManager.longitude

                // Gerar o timestamp
                val timestamp = System.currentTimeMillis()

                // Loga os valores de localização
                val messageLocation = "Localização na rotina - Latitude: $latitude, Longitude: $longitude"
                logWriter.writeLog(TAG, messageLocation)

                if (latitude != null && longitude != null) {
                    val json = JSONObject().apply {
                        put("timestamp", timestamp)
                        put("latitude", latitude)
                        put("longitude", longitude)
                    }
                    apiClient.postSendData(GPS_ENDPOINT, json.toString())
                    databaseManager.insertLocationData(latitude, longitude, timestamp, deviceId)
                }
            } catch (e: Exception) {
                logWriter.writeLog(TAG, "Erro ao buscar dados de localização: ${e.message}")
            }
        }
    }

    private suspend fun getDataCamera() {
        withContext(Dispatchers.IO) {
            try {
                // Captura uma imagem e a converte para Base64
                val base64Image = customCameraManager.initializeCameraAndTakePicture()
                logWriter.writeLog(TAG, "Imagem capturada em Base64: $base64Image")

                // Gerar o timestamp
                val timestamp = System.currentTimeMillis()

                if (base64Image != null) {
                    val json = JSONObject().apply {
                        put("timestamp", timestamp)
                        put("image_base64", base64Image)
                    }
                    apiClient.postSendData(PHOTO_ENDPOINT, json.toString())
                    databaseManager.insertImageData(base64Image, timestamp, deviceId)
                }
            } catch (e: Exception) {
                logWriter.writeLog(TAG, "Erro ao buscar dados da câmera: ${e.message}")
            }
        }
    }

    private fun getAndroidId(): String {
        return try {
            val androidId = Settings.Secure.getString(contentResolver, Settings.Secure.ANDROID_ID)
            logWriter.writeLog(TAG, "Identificador único do dispositivo: $androidId")
            androidId
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao obter identificador único do dispositivo: ${e.message}")
            "UnknownDevice"
        }
    }
}
