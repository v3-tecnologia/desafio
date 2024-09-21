package com.example.v3challenge.viewModel

import android.app.Application
import android.content.Context
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.DefaultLifecycleObserver
import androidx.lifecycle.LifecycleOwner
import androidx.lifecycle.ViewModel
import com.example.v3challenge.model.Gyro
import com.example.v3challenge.repository.LogsRepository
import com.mutualmobile.composesensors.GyroscopeSensorState
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.util.Timer
import javax.inject.Inject
import kotlin.concurrent.schedule

@HiltViewModel
class LogsViewModel @Inject constructor(
    private val application: Application,
    private val logsRepository: LogsRepository
) : ViewModel(), DefaultLifecycleObserver {
    private val context: Context by lazy { application.applicationContext }
    private val timer = Timer()
    private val TEN_SECONDS = 1000L
    private var faceDetected: Boolean = true
    private var gyro: MutableState<Gyro> = mutableStateOf(Gyro())
    var log: MutableState<String> = mutableStateOf("")


    fun startTimer() {
        timer.schedule(0L, TEN_SECONDS) {
            if(faceDetected) {

                sendGyroData()
                sendGpsData()
                sendPhoto()

                log.value += ""
            } else {
                log.value += "\nNo face detected!"
            }
        }
    }

    private fun sendGyroData() {
        CoroutineScope(Dispatchers.Main).launch {
            val result = "X: ${gyro.value.x}\nY: ${gyro.value.y}\nZ: ${gyro.value.z}"
            logsRepository.sendGyro(result)
        }
    }


    private fun sendGpsData() {
        CoroutineScope(Dispatchers.Main).launch {
            val result = ""
            logsRepository.sendGps(result)
        }
    }

    private fun sendPhoto() {
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
        gyro.value.x = event.xRotation.toString()
        gyro.value.y = event.yRotation.toString()
        gyro.value.z = event.zRotation.toString()
    }

}