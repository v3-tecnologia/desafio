package com.entregas.android_pleno_teste_v3.services

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Intent
import android.content.pm.ServiceInfo.FOREGROUND_SERVICE_TYPE_DATA_SYNC
import android.database.Cursor
import android.os.Build
import android.os.Handler
import android.os.IBinder
import android.provider.MediaStore
import android.util.Log
import android.widget.Toast
import androidx.core.app.NotificationCompat
import com.entregas.android_pleno_teste_v3.R
import java.io.File

class PhotoMonitorService : Service() {

    private val handler = Handler()
    private val checkInterval = 10000L // Intervalo de 10 segundos
    private lateinit var photoDirectory: File

    override fun onCreate() {
        super.onCreate()

        startForegroundService()
        setupPhotoDirectory()
        startMonitoringForNewPhotos()
    }

    private fun startForegroundService() {
        val notificationChannelId = "PHOTO_MONITOR_SERVICE_CHANNEL"

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val channel = NotificationChannel(
                notificationChannelId,
                "Photo Monitor Service",
                NotificationManager.IMPORTANCE_LOW
            )
            val manager = getSystemService(NotificationManager::class.java)
            manager.createNotificationChannel(channel)
        }

        val notification: Notification = NotificationCompat.Builder(this, notificationChannelId)
            .setContentTitle("Monitoring Photos")
            .setContentText("The service is running and monitoring photos.")
            .setSmallIcon(R.drawable.ic_launcher_foreground)
            .build()

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
            startForeground(1, notification,FOREGROUND_SERVICE_TYPE_DATA_SYNC)
        }else{
            startForeground(1, notification)

        }
    }

    private fun setupPhotoDirectory() {
        photoDirectory = getOutputDirectory()
    }

    private fun getOutputDirectory(): File {
        val mediaDir = externalMediaDirs.firstOrNull()?.let {
            File(it, resources.getString(R.string.app_name)).apply { mkdirs() }
        }
        return if (mediaDir != null && mediaDir.exists())
            mediaDir else filesDir
    }

    private fun startMonitoringForNewPhotos() {
        val runnable = object : Runnable {
            override fun run() {
                processLastPhoto()
                handler.postDelayed(this, checkInterval)
            }
        }
        handler.post(runnable)
    }

    private fun processLastPhoto() {
        val lastPhotoFile = getLastPhotoFile()
        lastPhotoFile?.let {
            // Process the last photo file
            Log.d("PhotoMonitorService", "Processing photo: ${it.absolutePath}")
            showToast("Last photo found: ${it.absolutePath}")
            // Add additional processing logic here if needed
        } ?: run {
            Log.d("PhotoMonitorService", "No photos found.")
            showToast("No photos found")
        }
    }

    private fun getLastPhotoFile(): File? {
        val projection = arrayOf(MediaStore.Images.Media._ID, MediaStore.Images.Media.DATA)
        val sortOrder = "${MediaStore.Images.Media.DATE_ADDED} DESC"
        val cursor: Cursor? = contentResolver.query(
            MediaStore.Images.Media.EXTERNAL_CONTENT_URI,
            projection,
            null,
            null,
            sortOrder
        )

        cursor?.use {
            if (it.moveToFirst()) {
                val dataIndex = it.getColumnIndexOrThrow(MediaStore.Images.Media.DATA)
                val photoPath = it.getString(dataIndex)
                return File(photoPath)
            }
        }
        return null
    }

    private fun showToast(message: String) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show()
    }

    override fun onDestroy() {
        super.onDestroy()
        handler.removeCallbacksAndMessages(null)
    }

    override fun onBind(intent: Intent?): IBinder? {
        return null
    }
}
