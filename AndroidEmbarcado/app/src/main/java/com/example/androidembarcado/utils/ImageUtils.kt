package com.example.androidembarcado.utils

import android.graphics.Bitmap
import android.graphics.Matrix

object ImageUtils {
    
    fun cropFaceFromBitmap(bitmap: Bitmap, x: Int, y: Int, width: Int, height: Int): Bitmap {
        return Bitmap.createBitmap(bitmap, x, y, width, height)
    }

    fun rotateBitmap(bitmap: Bitmap, rotation: Float): Bitmap {
        val matrix = Matrix()
        matrix.postRotate(rotation)
        return Bitmap.createBitmap(bitmap, 0, 0, bitmap.width, bitmap.height, matrix, true)
    }
}
