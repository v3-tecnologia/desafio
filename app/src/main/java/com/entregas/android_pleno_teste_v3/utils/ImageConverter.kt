package com.entregas.android_pleno_teste_v3.utils

import android.graphics.Bitmap
import android.graphics.BitmapFactory
import android.util.Base64
import java.io.ByteArrayOutputStream
import java.io.File

class ImageConverter {
    fun convertImageToBase64(imagePath: String): String? {
        val imageFile = File(imagePath)
        val bitmap = BitmapFactory.decodeFile(imageFile.absolutePath)

        return bitmap?.let {
            encodeBitmapToBase64(it)
        }
    }

    private fun encodeBitmapToBase64(bitmap: Bitmap): String {
        val byteArrayOutputStream = ByteArrayOutputStream()
        bitmap.compress(Bitmap.CompressFormat.PNG, 100, byteArrayOutputStream)
        val byteArray = byteArrayOutputStream.toByteArray()
        return Base64.encodeToString(byteArray, Base64.DEFAULT)
    }
}
