package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody.Companion.toRequestBody
import java.util.concurrent.TimeUnit

class ApiClient(private val context: Context) {

    private val logWriter: LogWriter = LogWriter(context)

    companion object {
        const val TAG = "ApiClient"
        private const val TIMEOUT = 10L // Timeout de 10 segundos
    }

    private val client: OkHttpClient = OkHttpClient.Builder()
        .connectTimeout(TIMEOUT, TimeUnit.SECONDS)
        .writeTimeout(TIMEOUT, TimeUnit.SECONDS)
        .readTimeout(30, TimeUnit.SECONDS)
        .build()

    fun postSendData(url: String, json: String): Boolean {
        return try {
            val requestBody = json.toRequestBody("application/json".toMediaTypeOrNull())
            val request = Request.Builder()
                .url(url)
                .post(requestBody)
                .build()

            client.newCall(request).execute().use { response ->
                if (response.isSuccessful) {
                    logWriter.writeLog(TAG, "Dados enviados com sucesso para $url")
                    true
                } else {
                    logWriter.writeLog(TAG, "Erro ao enviar dados para $url: ${response.message}")
                    false
                }
            }
        } catch (e: Exception) {
            logWriter.writeLog(TAG, "Erro ao enviar requisição para $url: ${e.message}")
            false
        }
    }
}
