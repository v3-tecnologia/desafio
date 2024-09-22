package com.example.v3challenge.viewModel

import android.app.Application
import android.content.Context
import android.util.Log
import androidx.collection.ArraySet
import androidx.collection.arraySetOf
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.DefaultLifecycleObserver
import androidx.lifecycle.LifecycleOwner
import androidx.lifecycle.ViewModel
import com.example.v3challenge.localData.PrefsInterface
import com.example.v3challenge.localData.PrefsRepository
import com.example.v3challenge.model.Gyro
import com.example.v3challenge.network.ApiSettings.TEN_SECONDS
import com.example.v3challenge.network.ApiSettings.moshi
import com.example.v3challenge.repository.LogsRepository
import com.mutualmobile.composesensors.GyroscopeSensorState
import com.squareup.moshi.JsonAdapter
import com.squareup.moshi.Moshi
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.util.Timer
import javax.inject.Inject
import kotlin.concurrent.schedule

@Suppress("UNCHECKED_CAST")
@HiltViewModel
class LogsViewModel @Inject constructor(
    private val application: Application,
    private val logsRepository: LogsRepository
) : ViewModel(), DefaultLifecycleObserver {
    private val context: Context by lazy { application.applicationContext }
    private val timer = Timer()
    private var faceDetected: Boolean = true
    private var currentGyroData: MutableState<Gyro> = mutableStateOf(Gyro())

    private val gyroAdapter: JsonAdapter<Set<*>>? = moshi.adapter(Set::class.java)

    private val gyroPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "gyro-data")
    }
    private val gpsPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "gps-data")
    }
    private val photoPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "photo-data")
    }

    var log: MutableState<String> = mutableStateOf("")

    private fun saveGyroDataLocally() {
        val gyroData = gyroPrefs.getPref()
        if (gyroData != null) {
            val gyroSet: LinkedHashSet<Gyro> = gyroAdapter?.fromJson(gyroData) as LinkedHashSet<Gyro>
            gyroSet.add(currentGyroData.value)
            gyroPrefs.setPref(gyroAdapter.toJson(gyroSet).toString())
        } else {
            gyroPrefs.setPref(gyroAdapter?.toJson(arraySetOf(currentGyroData.value)).toString())
        }
        Log.i("New gyro saved:", currentGyroData.value.toString())
    }

    fun startTimer() {
        timer.schedule(0L, TEN_SECONDS) {
            if(faceDetected) {
                saveAndSendGyroData()
                saveAndSendGpsData()
                saveAndSendPhotoData()
                log.value += ""
            } else {
                log.value += "\nNo face detected!"
            }
        }
    }

    private fun saveAndSendGyroData() {
        saveGyroDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            val result = "X: ${currentGyroData.value.x}\nY: ${currentGyroData.value.y}\nZ: ${currentGyroData.value.z}"
            logsRepository.sendGyro(result)
        }
    }

    private fun saveAndSendGpsData() {
        CoroutineScope(Dispatchers.Main).launch {
            val result = ""
            logsRepository.sendGps(result)
        }
    }

    private fun saveAndSendPhotoData() {
        CoroutineScope(Dispatchers.Main).launch {
            val result = ""
            logsRepository.sendPhoto(result)
        }
    }

    override fun onDestroy(owner: LifecycleOwner) {
        super.onDestroy(owner)
        timer.cancel()
    }

    fun setGyroData(event: GyroscopeSensorState) {
        currentGyroData.value.x = event.xRotation.toString()
        currentGyroData.value.y = event.yRotation.toString()
        currentGyroData.value.z = event.zRotation.toString()
        currentGyroData.value.timestamp = System.currentTimeMillis()
    }

}