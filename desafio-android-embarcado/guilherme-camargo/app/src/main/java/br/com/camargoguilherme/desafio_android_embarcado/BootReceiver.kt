package br.com.camargoguilherme.desafio_android_embarcado

import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.os.Build
import android.util.Log

class BootReceiver : BroadcastReceiver() {
    companion object {
        const val TAG = "BootReceiver"
    }
    private lateinit var logWriter: LogWriter
    override fun onReceive(context: Context, intent: Intent) {
        logWriter = LogWriter(context)
        logWriter.writeLog(TAG, "Aplicativo atualizado ou instalado")
        try {
            Log.i(BackgroundService.TAG, intent.action.toString())
            if (intent.action == Intent.ACTION_BOOT_COMPLETED) {
                val serviceIntent = Intent(context, BackgroundService::class.java)
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
                    logWriter.writeLog(AppUpdateReceiver.TAG,"Iniciando serviço como ForegroundService")
                    context.startForegroundService(serviceIntent)
                } else {
                    logWriter.writeLog(AppUpdateReceiver.TAG,"Iniciando serviço como BackgroundService")
                    context.startService(serviceIntent)
                }
            }
        }catch (e: Exception) {
            logWriter.writeLog(AppUpdateReceiver.TAG,"Erro ao iniciar serviço: ${e.message}")
        }

    }
}