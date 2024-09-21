package com.example.v3challenge.repository

import android.content.Context
import com.example.v3challenge.network.LogsInterface
import com.example.v3challenge.network.Resource
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.android.scopes.ActivityScoped
import retrofit2.HttpException
import java.lang.Exception
import javax.inject.Inject

@ActivityScoped
class LogsRepository @Inject constructor(
    private val logInterface: LogsInterface,
    @ApplicationContext val context: Context) {

    suspend fun sendGyro(data: String): Resource<String> {
        return try {
            val response = logInterface.sendGyro(data)
            Resource.Success(response)
        } catch (e: Exception) {
            //This should be the return if the service fails.
            //But I'm sending Success anyways because we don't have a real endpoint.
            //Resource.Error("Error sending Gyro data.", showError = false)

            Resource.Success("")
        }
    }

    suspend fun sendGps(data: String): Resource<String> {
        return try {
            Resource.Success(logInterface.sendGps(data))
        } catch (e: Exception) {
            Resource.Error("Error sending GPS data.", showError = false)
        }
    }

    suspend fun sendPhoto(data: String): Resource<String> {
        return try {
            Resource.Success(logInterface.sendPhoto(data))
        } catch (e: Exception) {
            Resource.Error("Error sending Photo data.", showError = false)
        }
    }

}