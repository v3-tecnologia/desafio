package br.com.camargoguilherme.desafio_android_embarcado

import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.util.Log

class AppUpdateReceiver : BroadcastReceiver() {
    companion object {
        const val TAG = "AppUpdateReceiver"
    }
    private lateinit var logWriter: LogWriter
    override fun onReceive(context: Context, intent: Intent) {
        logWriter = LogWriter(context)
        logWriter.writeLog(TAG, "Aplicativo atualizado ou instalado")

        try {
            // Inicia o serviço em background
            val serviceIntent = Intent(context, BackgroundService::class.java)
            if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.O) {
                logWriter.writeLog(TAG,"Iniciando serviço como ForegroundService")
                context.startForegroundService(serviceIntent)
            } else {
                logWriter.writeLog(TAG,"Iniciando serviço como BackService")
                context.startService(serviceIntent)
            }
        }catch (e: Exception){
            logWriter.writeLog(TAG,"Erro ao iniciar serviço: ${e.message}")
        }

    }
}