package com.example.v3challenge.viewModel

import android.app.Application
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.DefaultLifecycleObserver
import androidx.lifecycle.LifecycleOwner
import androidx.lifecycle.ViewModel
import com.example.v3challenge.repository.LogsRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import java.util.Timer
import javax.inject.Inject
import kotlin.concurrent.schedule

@HiltViewModel
class LogsViewModel @Inject constructor(
    private val application: Application,
    private val logsRepository: LogsRepository
) : ViewModel(), DefaultLifecycleObserver {

    private val timer = Timer()
    private val TEN_SECONDS = 1000L
    var log: MutableState<String> = mutableStateOf("")
    private var faceDetected: Boolean = false

    fun startTimer() {
        var counter = 0
        timer.schedule(0L, TEN_SECONDS) {
            if(faceDetected) {
                log.value += "\nLog"
            } else {
                log.value += "\nNo face detected!"
            }
        }
    }

    override fun onDestroy(owner: LifecycleOwner) {
        super.onDestroy(owner)
        timer.cancel()
    }

}