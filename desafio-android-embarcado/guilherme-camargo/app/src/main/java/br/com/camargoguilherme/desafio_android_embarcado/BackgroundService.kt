package br.com.camargoguilherme.desafio_android_embarcado

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Intent
import android.content.pm.PackageManager
import android.os.Build
import android.os.Handler
import android.os.IBinder
import android.os.Looper
import android.util.Log
import androidx.core.app.NotificationCompat
import androidx.core.content.ContextCompat

class BackgroundService : Service() {

    companion object {
        const val CHANNEL_ID = "BackgroundServiceChannel"
        const val TAG = "BackgroundService"
    }

    private lateinit var gyroscopeManager: GyroscopeManager
    private lateinit var locationManager: LocationManager
    private val handler = Handler(Looper.getMainLooper())
    private val interval: Long = 10000 // 10 segundos
    private val runnable = object : Runnable {
        override fun run() {
            executeRoutine()
            handler.postDelayed(this, interval)
        }
    }

    override fun onCreate() {
        super.onCreate()
        Log("Service criado")

        // Inicializa o GyroscopeManager e LocationManager
        gyroscopeManager = GyroscopeManager(this)
        locationManager = LocationManager(this)

        // Certifique-se de que as permissões de localização foram concedidas antes de iniciar as atualizações
        if (checkPermissions()) {
            locationManager.startLocationUpdates()
        } else {
            Log("Permissões de localização não concedidas")
        }

        // Configura o canal de notificação e inicia o serviço em foreground
        createNotificationChannel()
        val notification: Notification = NotificationCompat
            .Builder(this, CHANNEL_ID)
            .setContentTitle("Background Service")
            .setContentText("Service is running in the background")
            //.setSmallIcon(androidx.core.R.drawable.notification_action_background)
            .setPriority(NotificationCompat.PRIORITY_HIGH)
            .build()

        startForeground(1, notification)

        // Inicia a rotina imediatamente
        handler.post(runnable)
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        Log("Serviço iniciado")
        return START_STICKY
    }

    override fun onDestroy() {
        super.onDestroy()
        Log("Serviço destruído")

        // Para as atualizações de localização e remove as callbacks da rotina
        locationManager.stopLocationUpdates()
        handler.removeCallbacks(runnable)
    }

    override fun onBind(intent: Intent?): IBinder? {
        return null
    }

    private fun createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val serviceChannel = NotificationChannel(
                CHANNEL_ID,
                "Foreground Service Channel",
                NotificationManager.IMPORTANCE_LOW
            )
            val manager = getSystemService(NotificationManager::class.java)
            manager.createNotificationChannel(serviceChannel)
        }
    }

    private fun checkPermissions(): Boolean {
        // Verifica se as permissões de localização foram concedidas
        return ContextCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_FINE_LOCATION) == PackageManager.PERMISSION_GRANTED &&
                ContextCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_COARSE_LOCATION) == PackageManager.PERMISSION_GRANTED
    }

    private fun executeRoutine() {
        // Acessa os valores do giroscópio
        val x = gyroscopeManager.xValue
        val y = gyroscopeManager.yValue
        val z = gyroscopeManager.zValue

        // Acessa os valores de localização
        val latitude = locationManager.latitude
        val longitude = locationManager.longitude

        // Loga os valores do giroscópio
        val messageGyroscope = "Giroscópio na rotina - X: $x, Y: $y, Z: $z"
        Log(messageGyroscope)

        // Loga os valores de localização
        val messageLocation = "Localização na rotina - Latitude: $latitude, Longitude: $longitude"
        Log(messageLocation)
    }

    private fun Log(message: String) {
        Log.i(TAG, message)
        Log.e(TAG, message)
        Log.d(TAG, message)
        Log.v(TAG, message)
    }
}
