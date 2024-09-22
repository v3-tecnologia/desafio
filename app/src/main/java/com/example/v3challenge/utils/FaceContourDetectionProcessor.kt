package com.example.v3challenge.utils

import android.media.Image
import android.util.Log
import com.google.android.gms.tasks.Task
import com.google.mlkit.vision.common.InputImage
import com.google.mlkit.vision.face.Face
import com.google.mlkit.vision.face.FaceDetection
import com.google.mlkit.vision.face.FaceDetectorOptions
import java.io.IOException

class FaceContourDetectionProcessor(
    private val onSuccessCallback: (FaceStatus, Image?) -> Unit) :
    BaseImageAnalyzer<List<Face>>() {

    private val realTimeOpts = FaceDetectorOptions.Builder()
        .setPerformanceMode(FaceDetectorOptions.PERFORMANCE_MODE_FAST)
        .setContourMode(FaceDetectorOptions.CONTOUR_MODE_NONE)
        .build()

    private val detector = FaceDetection.getClient(realTimeOpts)

    override fun detectInImage(image: InputImage): Task<List<Face>> {
        return detector.process(image)
    }

    override fun stop() {
        try {
            detector.close()
        } catch (e: IOException) {
            Log.e(TAG, "Exception thrown while trying to close Face Detector: $e")
        }
    }

    override fun onSuccess(
        results: List<Face>,
        image: Image,
    ) {
        if (results.isNotEmpty()) {
            onSuccessCallback(FaceStatus.VALID, image)
        } else {
            onSuccessCallback(FaceStatus.NO_FACE, null)
        }
    }

    override fun onFailure(e: Exception) {
        Log.e(TAG, "Face Detector failed. $e")
        onSuccessCallback(FaceStatus.NO_FACE, null)
    }

    companion object {
        private const val TAG = "FaceDetectorProcessor"
    }

}