package com.entregas.android_pleno_teste_v3.broadcast

import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.os.Build
import com.entregas.android_pleno_teste_v3.PermissionRequestActivity
import com.entregas.android_pleno_teste_v3.services.BackgroundService
import com.entregas.android_pleno_teste_v3.services.PhotoMonitorService
import com.entregas.android_pleno_teste_v3.utils.PermissionUtils

class StartupReceiver : BroadcastReceiver() {

    override fun onReceive(context: Context, intent: Intent) {
        if (hasPermissions(context)) {
            startService(context)
        } else {
            requestPermissions(context)
        }
    }

    private fun hasPermissions(context: Context): Boolean {
        // Verifica se as permissões estão concedidas
        return PermissionUtils.hasRequiredPermissions(context)
    }

    private fun requestPermissions(context: Context) {
        // Inicia a Activity para solicitar permissões
        val intent = Intent(context, PermissionRequestActivity::class.java)
        intent.flags = Intent.FLAG_ACTIVITY_NEW_TASK
        context.startActivity(intent)
    }

     fun startService(context: Context) {
        val serviceIntent = Intent(context, BackgroundService::class.java)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            context.startForegroundService(serviceIntent)
        } else {
            context.startService(serviceIntent)
        }

        val photoServiceIntent = Intent(context, PhotoMonitorService::class.java)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            context.startForegroundService(photoServiceIntent)
        } else {
            context.startService(photoServiceIntent)
        }
    }
}
