package com.example.v3challenge.viewModel

import android.Manifest
import android.annotation.SuppressLint
import android.app.Application
import android.content.Context
import android.content.pm.PackageManager
import android.util.Log
import androidx.collection.arraySetOf
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import androidx.core.app.ActivityCompat
import androidx.lifecycle.DefaultLifecycleObserver
import androidx.lifecycle.LifecycleOwner
import androidx.lifecycle.ViewModel
import com.example.v3challenge.localData.PrefsInterface
import com.example.v3challenge.localData.PrefsRepository
import com.example.v3challenge.model.Gps
import com.example.v3challenge.model.Gyro
import com.example.v3challenge.network.ApiSettings.TEN_SECONDS
import com.example.v3challenge.network.ApiSettings.moshi
import com.example.v3challenge.repository.LogsRepository
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationServices
import com.mutualmobile.composesensors.GyroscopeSensorState
import com.squareup.moshi.JsonAdapter
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
    private var currentGpsData: MutableState<Gps> = mutableStateOf(Gps())

    private val setAdapter: JsonAdapter<Set<*>>? = moshi.adapter(Set::class.java)

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

    private var fusedLocationProviderClient: FusedLocationProviderClient = LocationServices.getFusedLocationProviderClient(context)

    //Start Functions
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

    private fun saveGyroDataLocally() {
        val gyroData = gyroPrefs.getPref()
        if (gyroData != null) {
            val gyroSet: LinkedHashSet<Gyro> = setAdapter?.fromJson(gyroData) as LinkedHashSet<Gyro>
            gyroSet.add(currentGyroData.value)
            gyroPrefs.setPref(setAdapter.toJson(gyroSet).toString())
        } else {
            gyroPrefs.setPref(setAdapter?.toJson(arraySetOf(currentGyroData.value)).toString())
        }
//        Log.i("New gyro saved:", currentGyroData.value.toString())
    }

    @SuppressLint("MissingPermission")
    private fun saveGpsDataLocally() {

        // Retrieve the last known location
        fusedLocationProviderClient.lastLocation
            .addOnSuccessListener { location ->
                location?.let {
                    currentGpsData.value.lat = it.latitude
                    currentGpsData.value.lon = it.longitude
                    currentGpsData.value.timestamp = System.currentTimeMillis()
                }
            }
            .addOnFailureListener { exception ->
                // If an error occurs, invoke the failure callback with the exception
                // TODO
            }

        val gpsData = gpsPrefs.getPref()
        if (gpsData != null) {
            val gpsSet: LinkedHashSet<Gps> = setAdapter?.fromJson(gpsData) as LinkedHashSet<Gps>
            gpsSet.add(currentGpsData.value)
            gyroPrefs.setPref(setAdapter.toJson(gpsSet).toString())
        } else {
            gyroPrefs.setPref(setAdapter?.toJson(arraySetOf(currentGpsData.value)).toString())
        }
        Log.i("New GPS saved:", currentGpsData.value.toString())
    }

    private fun saveAndSendGyroData() {
        saveGyroDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            logsRepository.sendGyro(currentGyroData.value.toString())
        }
    }

    private fun saveAndSendGpsData() {
        if (hasNotLocationPermissions()) return

        saveGpsDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            val result = ""
            logsRepository.sendGps(result)
        }
    }

    private fun hasNotLocationPermissions(): Boolean {
        return ActivityCompat.checkSelfPermission(
            context,
            Manifest.permission.ACCESS_FINE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(
            context,
            Manifest.permission.ACCESS_COARSE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED
    }

    private fun saveAndSendPhotoData() {
        //savePhotoDataLocally()
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