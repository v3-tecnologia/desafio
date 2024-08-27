package com.entregas.android_pleno_teste_v3.services

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.Service
import android.content.Intent
import android.content.pm.ServiceInfo
import android.graphics.BitmapFactory
import android.os.Build
import android.os.Environment
import android.os.Handler
import android.os.IBinder
import android.util.Log
import android.widget.Toast
import androidx.core.app.NotificationCompat
import com.entregas.android_pleno_teste_v3.R
import com.entregas.android_pleno_teste_v3.data.PhotoEntity
import com.entregas.android_pleno_teste_v3.data.database.AppDatabase
import com.entregas.android_pleno_teste_v3.domain.PhotoRequestDataClass
import com.entregas.android_pleno_teste_v3.repository.FotoRequestRepository
import com.entregas.android_pleno_teste_v3.utils.Constants.Companion.CAPTURE_INTERVAL
import com.entregas.android_pleno_teste_v3.utils.GetMacAddress
import com.entregas.android_pleno_teste_v3.utils.ImageConverter
import com.google.mlkit.vision.face.FaceDetection
import com.google.mlkit.vision.face.FaceDetector
import com.google.mlkit.vision.face.FaceDetectorOptions
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.io.File

class PhotoMonitorService : Service() {

    private val handler = Handler()
    private val checkInterval = CAPTURE_INTERVAL
    private var lastProcessedFile: File? = null
    private lateinit var dataBase: AppDatabase
    private lateinit var  coroutineScope: CoroutineScope
    private val fotoRequestRepository = FotoRequestRepository()

    override fun onCreate() {
        super.onCreate()
        startForegroundService()
        startMonitoringForNewPhotos()
        dataBase = AppDatabase.getDatabase(applicationContext)
        coroutineScope = CoroutineScope(Dispatchers.IO)
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
            startForeground(1, notification, ServiceInfo.FOREGROUND_SERVICE_TYPE_DATA_SYNC)
        } else {
            startForeground(1, notification)
        }
    }

    private fun startMonitoringForNewPhotos() {
        val runnable = object : Runnable {
            override fun run() {
                processNewPhotos()
                handler.postDelayed(this, checkInterval)
            }
        }
        handler.post(runnable)
    }

    private fun processNewPhotos() {
        val newPhotoFile = getLastPhotoFromGallery()
        newPhotoFile?.let {
            Log.d("PhotoMonitorService", "Processing new photo: ${it.absolutePath}")

            checkForFacesInPhoto(it)
            lastProcessedFile = it
        } ?: run {
            Log.d("PhotoMonitorService", "No new photos found.")

        }
    }

    private fun getLastPhotoFromGallery(): File? {
        val directory =
            Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DCIM).toString()
        val photoDirectory = File(directory)
        if (photoDirectory.exists()) {
            val files = photoDirectory.listFiles { _, name ->
                name.endsWith(".jpg") || name.endsWith(".jpeg") || name.endsWith(".png")
            }
            return files?.maxByOrNull { it.lastModified() }
        }
        return null
    }

    private fun checkForFacesInPhoto(photoFile: File) {
        val bitmap = BitmapFactory.decodeFile(photoFile.absolutePath)

        val options = FaceDetectorOptions.Builder()
            .setPerformanceMode(FaceDetectorOptions.PERFORMANCE_MODE_FAST)
            .setLandmarkMode(FaceDetectorOptions.LANDMARK_MODE_NONE)
            .setClassificationMode(FaceDetectorOptions.CLASSIFICATION_MODE_NONE)
            .build()

        val detector: FaceDetector = FaceDetection.getClient(options)
        val image = com.google.mlkit.vision.common.InputImage.fromBitmap(bitmap, 0)

        detector.process(image)
            .addOnSuccessListener { faces ->
                if (faces.isNotEmpty()) {
                    coroutineScope.launch {
                        dataBase.photoScopeDao().insert(
                            PhotoEntity(
                                2,
                                ImageConverter.convertImageToBase64(bitmap),
                            )
                        )
                        fotoRequestRepository.sendPhotoRequest(
                            PhotoRequestDataClass(
                                ImageConverter.convertImageToBase64(bitmap),
                                GetMacAddress(applicationContext).getUniqueDeviceId()
                            )
                        ).collect{
                            if(it.isSuccess){

                            }else{

                            }
                        }
                    }
                    Log.d("PhotoMonitorService", "Faces detected: ${faces.size}")
                } else {

                    Log.d("PhotoMonitorService", "No faces detected.")
                }
            }
            .addOnFailureListener { e ->
                showToast("Face detection failed.")
                Log.e("PhotoMonitorService", "Face detection failed", e)
            }
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
