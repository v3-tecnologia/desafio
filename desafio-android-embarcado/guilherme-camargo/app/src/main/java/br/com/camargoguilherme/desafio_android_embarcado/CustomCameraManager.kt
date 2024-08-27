package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.graphics.*
import android.hardware.camera2.*
import android.media.ImageReader
import android.util.Base64
import com.google.mlkit.vision.common.InputImage
import com.google.mlkit.vision.face.Face
import com.google.mlkit.vision.face.FaceDetection
import com.google.mlkit.vision.face.FaceDetectorOptions
import java.io.ByteArrayOutputStream
import java.io.IOException

class CustomCameraManager(private val context: Context) {

    companion object {
        const val TAG = "CustomCameraManager"
    }

    private var cameraDevice: CameraDevice? = null
    private var imageReader: ImageReader? = null
    private val logWriter = LogWriter(context)

    fun initializeCameraAndTakePicture(): String? {
        var base64Image: String? = null
        try {
            val cameraManager = context.getSystemService(Context.CAMERA_SERVICE) as CameraManager
            val cameraId = cameraManager.cameraIdList[1] // Assumindo que a primeira câmera é a traseira

            val cameraCharacteristics = cameraManager.getCameraCharacteristics(cameraId)
            val streamConfigurationMap = cameraCharacteristics.get(CameraCharacteristics.SCALER_STREAM_CONFIGURATION_MAP)
            val imageDimension = streamConfigurationMap?.getOutputSizes(ImageFormat.JPEG)?.get(0)

            imageReader = ImageReader.newInstance(imageDimension!!.width, imageDimension.height, ImageFormat.JPEG, 1)

            cameraManager.openCamera(cameraId, object : CameraDevice.StateCallback() {
                override fun onOpened(camera: CameraDevice) {
                    cameraDevice = camera
                    logWriter.writeLog(TAG, "Câmera aberta com sucesso")
                    base64Image = captureImage()
                }

                override fun onDisconnected(camera: CameraDevice) {
                    cameraDevice?.close()
                    cameraDevice = null
                    logWriter.writeLog(TAG, "Câmera desconectada")
                }

                override fun onError(camera: CameraDevice, error: Int) {
                    cameraDevice?.close()
                    cameraDevice = null
                    when (error) {
                        ERROR_CAMERA_DISABLED -> {
                            logWriter.writeLog(TAG, "Erro ao abrir a câmera: Câmera desativada pelo sistema (Código de erro $error)")
                        }
                        ERROR_CAMERA_DEVICE -> {
                            logWriter.writeLog(TAG, "Erro ao abrir a câmera: Falha do dispositivo (Código de erro $error)")
                        }
                        ERROR_CAMERA_SERVICE -> {
                            logWriter.writeLog(TAG, "Erro ao abrir a câmera: Serviço de câmera inacessível (Código de erro $error)")
                        }
                        ERROR_CAMERA_IN_USE -> {
                            logWriter.writeLog(TAG, "Erro ao abrir a câmera: Câmera já em uso (Código de erro $error)")
                        }
                        else -> {
                            logWriter.writeLog(TAG, "Erro desconhecido ao abrir a câmera (Código de erro $error)")
                        }
                    }
                }
            }, null)

        } catch (e: CameraAccessException) {
            logWriter.writeLog(TAG, "Erro de acesso à câmera: ${e.message}")
        } catch (e: SecurityException) {
            logWriter.writeLog(TAG, "Erro de segurança ao acessar a câmera: ${e.message}")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro desconhecido ao iniciar a câmera: ${e.message}")
        }

        return base64Image
    }

    private fun captureImage(): String? {
        var base64Image: String? = null
        try {
            if (cameraDevice == null) {
                logWriter.writeLog(TAG, "Câmera não está inicializada")
                return null
            }

            val captureRequestBuilder = cameraDevice?.createCaptureRequest(CameraDevice.TEMPLATE_STILL_CAPTURE)
            captureRequestBuilder?.addTarget(imageReader?.surface!!)

            cameraDevice?.createCaptureSession(listOf(imageReader?.surface), object : CameraCaptureSession.StateCallback() {
                override fun onConfigured(session: CameraCaptureSession) {
                    try {
                        session.capture(captureRequestBuilder!!.build(), object : CameraCaptureSession.CaptureCallback() {
                            override fun onCaptureCompleted(session: CameraCaptureSession, request: CaptureRequest, result: TotalCaptureResult) {
                                logWriter.writeLog(TAG, "Foto capturada com sucesso")
                                base64Image = processImageForFaceDetection()
                            }

                            override fun onCaptureFailed(session: CameraCaptureSession, request: CaptureRequest, failure: CaptureFailure) {
                                logWriter.writeLog(TAG, "Falha ao capturar foto: ${failure.reason}")
                            }
                        }, null)
                    } catch (e: CameraAccessException) {
                        logWriter.writeLog(TAG, "Erro ao capturar imagem: ${e.message}")
                    }
                }

                override fun onConfigureFailed(session: CameraCaptureSession) {
                    logWriter.writeLog(TAG, "Falha na configuração da sessão de captura")
                }
            }, null)

        } catch (e: CameraAccessException) {
            logWriter.writeLog(TAG, "Erro de acesso à câmera ao capturar foto: ${e.message}")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro desconhecido ao capturar foto: ${e.message}")
        }

        return base64Image
    }

    private fun processImageForFaceDetection(): String? {
        val image = imageReader?.acquireLatestImage() ?: return null

        val buffer = image.planes[0].buffer
        val bytes = ByteArray(buffer.remaining())
        buffer.get(bytes)

        val bitmap = BitmapFactory.decodeByteArray(bytes, 0, bytes.size)

        // Configura o detector de rostos
        val options = FaceDetectorOptions.Builder()
            .setPerformanceMode(FaceDetectorOptions.PERFORMANCE_MODE_FAST)
            .setLandmarkMode(FaceDetectorOptions.LANDMARK_MODE_NONE)
            .setClassificationMode(FaceDetectorOptions.CLASSIFICATION_MODE_NONE)
            .build()

        val detector = FaceDetection.getClient(options)
        val inputImage = InputImage.fromBitmap(bitmap, 0)

        var base64Image: String? = null

        detector.process(inputImage)
            .addOnSuccessListener { faces ->
                if (faces.isNotEmpty()) {
                    val croppedBitmap = cropFaceFromImage(bitmap, faces[0])
                    base64Image = bitmapToBase64(croppedBitmap)
                    logWriter.writeLog(TAG, "Rosto detectado e imagem processada")
                } else {
                    logWriter.writeLog(TAG, "Nenhum rosto detectado na imagem")
                }
            }
            .addOnFailureListener { e ->
                logWriter.writeLog(TAG, "Erro ao processar imagem para detecção de rosto: ${e.message}")
            }

        image.close()
        return base64Image
    }

    private fun cropFaceFromImage(bitmap: Bitmap, face: Face): Bitmap {
        val bounds = face.boundingBox
        return Bitmap.createBitmap(
            bitmap,
            bounds.left.coerceAtLeast(0),
            bounds.top.coerceAtLeast(0),
            bounds.width().coerceAtMost(bitmap.width - bounds.left),
            bounds.height().coerceAtMost(bitmap.height - bounds.top)
        )
    }

    private fun bitmapToBase64(bitmap: Bitmap): String? {
        return try {
            val byteArrayOutputStream = ByteArrayOutputStream()
            bitmap.compress(Bitmap.CompressFormat.JPEG, 100, byteArrayOutputStream)
            val byteArray = byteArrayOutputStream.toByteArray()
            Base64.encodeToString(byteArray, Base64.DEFAULT)
        } catch (e: IOException) {
            logWriter.writeLog(TAG, "Erro ao converter a imagem para Base64: ${e.message}")
            null
        }
    }

    fun stopCamera() {
        try {
            cameraDevice?.close()
            imageReader?.close()
            logWriter.writeLog(TAG, "Câmera parada com sucesso")
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao parar a câmera: ${e.message}")
        }
    }
}
