package com.entregas.android_pleno_teste_v3

import android.Manifest
import android.content.Intent
import android.content.pm.PackageManager
import android.os.Build
import android.os.Bundle
import android.provider.Settings
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import com.entregas.android_pleno_teste_v3.services.BackgroundService

class MainActivity : AppCompatActivity() {

    private val requestCodePermission = 1001

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        if (hasLocationPermission()) {
            startServiceAndFinish()
        } else {
            requestLocationPermission()
        }
    }
    private fun hasLocationPermission(): Boolean {
        return ActivityCompat.checkSelfPermission(
            this,
            Manifest.permission.ACCESS_FINE_LOCATION
        ) == PackageManager.PERMISSION_GRANTED || ActivityCompat.checkSelfPermission(
            this,
            Manifest.permission.ACCESS_COARSE_LOCATION
        ) == PackageManager.PERMISSION_GRANTED
    }

    private fun requestLocationPermission() {
        val intent = Intent(Settings.ACTION_APPLICATION_DETAILS_SETTINGS)
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
        val uri = android.net.Uri.fromParts("package", packageName, null)
        intent.data = uri
        startActivityForResult(intent, requestCodePermission)
    }

    private fun startServiceAndFinish() {
        val serviceIntent = Intent(this, BackgroundService::class.java)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            startForegroundService(serviceIntent)
        } else {
            startService(serviceIntent)
        }
        }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (requestCode == requestCodePermission) {
            if (hasLocationPermission()) {
                startServiceAndFinish()
            } else {
                // Optional: Notify user to grant permission if not granted
                Log.d("MainActivity", "Location permission not granted")
            }
        }
    }
}
