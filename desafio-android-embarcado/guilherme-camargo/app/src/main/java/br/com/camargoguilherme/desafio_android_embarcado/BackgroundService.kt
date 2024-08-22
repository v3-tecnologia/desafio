package br.com.camargoguilherme.desafio_android_embarcado

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Intent
import android.os.Build
import android.os.Handler
import android.os.IBinder
import android.os.Looper
import androidx.core.app.NotificationCompat

class BackgroundService : Service() {

    companion object {
        const val CHANNEL_ID = "BackgroundServiceChannel"
        const val TAG = "BackgroundService"
    }

    private lateinit var logWriter: LogWriter
    private lateinit var gyroscopeManager: GyroscopeManager
    private lateinit var locationManager: LocationManager
    private lateinit var customCameraManager: CustomCameraManager
    private val handler = Handler(Looper.getMainLooper())
    private val interval: Long = 10000 // 10 segundos
    private val runnable = object : Runnable {
        override fun run() {
            try {
                executeRoutine()
            } catch (e: Exception) {
                logWriter.writeLog(TAG, "Erro na execução da rotina: ${e.message}")
            }
            handler.postDelayed(this, interval)
        }
    }

    override fun onCreate() {
        super.onCreate()

        logWriter = LogWriter(this)
        logWriter.writeLog(TAG, "Service criado")

        // Inicia o serviço em foreground o mais rápido possível
        // Comentado pois o dispositivo matava a execução
        //startForegroundServiceWithNotification()

        // Inicializa o GyroscopeManager ou LocationManager ou CustomCameraManager
        try {
            gyroscopeManager = GyroscopeManager(this)
            locationManager = LocationManager(this, 1000, 500f)
            customCameraManager = CustomCameraManager(this)
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao inicializar GyroscopeManager ou LocationManager ou CustomCameraManager: ${e.message}")
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

    private fun startForegroundServiceWithNotification() {
        try {
            createNotificationChannel()
            val notification: Notification = NotificationCompat.Builder(this, CHANNEL_ID)
                .setContentTitle("Background Service")
                .setContentText("Service is running in the background")
                //.setSmallIcon(androidx.core.R.drawable.notification_action_background)
                .setPriority(NotificationCompat.PRIORITY_HIGH)
                .build()

            startForeground(1, notification)
            logWriter.writeLog(TAG, "Serviço movido para foreground com notificação")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao iniciar serviço em foreground: ${e.message}")
        }
    }

    private fun createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            try {
                val serviceChannel = NotificationChannel(
                    CHANNEL_ID,
                    "Foreground Service Channel",
                    NotificationManager.IMPORTANCE_LOW
                )
                val manager = getSystemService(NotificationManager::class.java)
                manager.createNotificationChannel(serviceChannel)
            } catch (e: Exception) {
                logWriter.writeLog(TAG, "Erro ao criar o canal de notificação: ${e.message}")
                throw e
            }
        }
    }

    private fun executeRoutine() {
        try {
            // Acessa os valores do giroscópio
            val x = gyroscopeManager.xValue
            val y = gyroscopeManager.yValue
            val z = gyroscopeManager.zValue

            // Acessa os valores de localização
            val latitude = locationManager.latitude
            val longitude = locationManager.longitude

            // Loga os valores do giroscópio
            val messageGyroscope = "Giroscópio na rotina - X: $x, Y: $y, Z: $z"
            logWriter.writeLog(TAG, messageGyroscope)

            // Loga os valores de localização
            val messageLocation = "Localização na rotina - Latitude: $latitude, Longitude: $longitude"
            logWriter.writeLog(TAG, messageLocation)

            // Captura uma imagem e a converte para Base64
            val base64Image = customCameraManager.initializeCameraAndTakePicture()
            logWriter.writeLog(TAG, "Imagem capturada em Base64: $base64Image")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao executar rotina: ${e.message}")
        }
    }
}
