package com.example.v3challenge.repository

import android.annotation.SuppressLint
import android.content.Context
import android.net.wifi.WifiManager
import com.example.v3challenge.network.LogsInterface
import com.example.v3challenge.network.Resource
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.android.scopes.ActivityScoped
import javax.inject.Inject


@ActivityScoped
class LogsRepository @Inject constructor(
    private val logInterface: LogsInterface,
    @ApplicationContext val context: Context
) {

    private var wifiManager: WifiManager =
        context.getSystemService(Context.WIFI_SERVICE) as WifiManager

    @SuppressLint("HardwareIds")
    val macAddress: String = wifiManager.connectionInfo.macAddress

    suspend fun sendGyro(data: String): Resource<String> {
        return try {
            val response = logInterface.sendGyro(data, macAddress)
            Resource.Success(response)
        } catch (e: Exception) {
            //TODO
            // Should add error treatments for each service
            Resource.Error("Error sending Gyro data.", showError = false)
        }
    }

    suspend fun sendGps(data: String): Resource<String> {
        return try {
            Resource.Success(logInterface.sendGps(data, macAddress))
        } catch (e: Exception) {
            Resource.Error("Error sending GPS data.", showError = false)
        }
    }

    suspend fun sendPhoto(data: String): Resource<String> {
        return try {
            Resource.Success(logInterface.sendPhoto(data, macAddress))
        } catch (e: Exception) {
            Resource.Error("Error sending Photo data.", showError = false)
        }
    }

}