package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.graphics.ImageFormat
import android.hardware.camera2.*
import android.media.ImageReader
import android.util.Base64
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

            logWriter.writeLog(TAG, "Total de cameras no dispositivo: ${cameraManager.cameraIdList.size}")

            for(cameraId: String in cameraManager.cameraIdList) {
                logWriter.writeLog(TAG, "Camera Id: $cameraId")
            }

            val cameraCharacteristics = cameraManager.getCameraCharacteristics(cameraId)
            val streamConfigurationMap = cameraCharacteristics.get(CameraCharacteristics.SCALER_STREAM_CONFIGURATION_MAP)
            val imageDimension = streamConfigurationMap?.getOutputSizes(ImageFormat.JPEG)?.get(0)

            imageReader = ImageReader.newInstance(imageDimension!!.width, imageDimension.height, ImageFormat.JPEG, 1)

            cameraManager.openCamera(cameraId, object : CameraDevice.StateCallback() {
                override fun onOpened(camera: CameraDevice) {
                    cameraDevice = camera
                    logWriter.writeLog(TAG, "Câmera aberta com sucesso")
                    // Iniciar a captura de imagem
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
                            // Mostrar uma mensagem apropriada ao usuário
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
                                base64Image = saveImageAsBase64()
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

    private fun saveImageAsBase64(): String? {
        val image = imageReader?.acquireLatestImage() ?: run {
            logWriter.writeLog(TAG, "Nenhuma imagem foi capturada")
            return null
        }

        return try {
            val buffer = image.planes[0].buffer
            val bytes = ByteArray(buffer.remaining())
            buffer.get(bytes)

            // Converte a imagem em Base64
            val outputStream = ByteArrayOutputStream()
            outputStream.write(bytes)
            val base64Image = Base64.encodeToString(outputStream.toByteArray(), Base64.DEFAULT)

            logWriter.writeLog(TAG, "Imagem convertida para Base64 com sucesso")
            base64Image

        } catch (e: IOException) {
            logWriter.writeLog(TAG, "Erro ao converter a imagem para Base64: ${e.message}")
            null
        } finally {
            image.close()
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
