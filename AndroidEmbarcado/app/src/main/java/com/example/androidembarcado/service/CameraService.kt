package com.example.androidembarcado.service

import android.content.Context
import android.graphics.Bitmap
import android.os.Environment
import com.example.androidembarcado.model.PhotoData
import com.example.androidembarcado.repository.TelemetryRepository
import com.example.androidembarcado.utils.ImageUtils
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.io.File
import java.io.FileOutputStream

class CameraService(private val context: Context) {

    private val repository: TelemetryRepository = TelemetryRepository(context)

    fun capturePhoto(bitmap: Bitmap, deviceId: String) {
        val file = savePhoto(bitmap)

        val photoData = PhotoData(
            photoPath = file.absolutePath,
            timestamp = System.currentTimeMillis(),
            deviceId = deviceId
        )

        CoroutineScope(Dispatchers.IO).launch {
            repository.insertPhotoData(photoData)
        }
    }

    private fun savePhoto(bitmap: Bitmap): File {
        val photoDir = File(context.getExternalFilesDir(Environment.DIRECTORY_PICTURES), "photos")
        if (!photoDir.exists()) photoDir.mkdirs()

        val fileName = "photo_${System.currentTimeMillis()}.jpg"
        val photoFile = File(photoDir, fileName)

        FileOutputStream(photoFile).use { fos ->
            bitmap.compress(Bitmap.CompressFormat.JPEG, 100, fos)
            fos.flush()
        }

        return photoFile
    }
}
