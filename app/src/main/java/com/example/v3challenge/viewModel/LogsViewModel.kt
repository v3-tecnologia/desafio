package com.example.v3challenge.viewModel

import android.Manifest
import android.annotation.SuppressLint
import android.app.Application
import android.content.Context
import android.content.pm.PackageManager
import android.graphics.Bitmap
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
import com.example.v3challenge.model.GenericLog
import com.example.v3challenge.model.Gps
import com.example.v3challenge.model.Gyro
import com.example.v3challenge.model.Photo
import com.example.v3challenge.network.ApiSettings.TEN_SECONDS
import com.example.v3challenge.network.ApiSettings.moshi
import com.example.v3challenge.repository.LogsRepository
import com.example.v3challenge.utils.FaceStatus
import com.example.v3challenge.utils.LogType
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationServices
import com.mutualmobile.composesensors.GyroscopeSensorState
import com.squareup.moshi.JsonAdapter
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.util.ArrayList
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
    private var currentGyroData: MutableState<Gyro> = mutableStateOf(Gyro())
    private var currentGpsData: MutableState<Gps> = mutableStateOf(Gps())
    private var currentPhotoData: MutableState<Photo> = mutableStateOf(Photo())
    private var isFaceDetectedNow: MutableState<Boolean> = mutableStateOf(false)
    private val setAdapter: JsonAdapter<Set<*>>? = moshi.adapter(Set::class.java)
    private var fusedLocationProviderClient: FusedLocationProviderClient =
        LocationServices.getFusedLocationProviderClient(context)

    private val gyroPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "gyro-data")
    }
    private val gpsPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "gps-data")
    }
    private val photoPrefs: PrefsInterface by lazy {
        PrefsRepository(context, "photo-data")
    }

    val logs: MutableList<GenericLog> = mutableListOf()

    //Start Functions
    fun startTimer() {
        timer.schedule(0L, TEN_SECONDS) {
            saveAndSendGyroData()
            saveAndSendGpsData()
            if (isFaceDetectedNow.value) {
                saveAndSendPhotoData()
            } else {
                currentPhotoData.value.photo = null
                currentPhotoData.value.timestamp = null
                Log.e("photo","No Face detected.")
            }
            logs.add(GenericLog(LogType.PHOTO, currentPhotoData.value.copy()))
        }
    }

    private fun hasNoLocationPermissions(): Boolean {
        return ActivityCompat.checkSelfPermission(
            context,
            Manifest.permission.ACCESS_FINE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(
            context,
            Manifest.permission.ACCESS_COARSE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED
    }

    private fun saveGyroDataLocally() {
        val gyroData = gyroPrefs.getPref()
        if (gyroData != null) {
            val gyroSet: LinkedHashSet<Gyro> = setAdapter?.fromJson(gyroData) as LinkedHashSet<Gyro>
            gyroSet.add(currentGyroData.value.copy())
            gyroPrefs.setPref(setAdapter.toJson(gyroSet))
        } else {
            gyroPrefs.setPref(setAdapter?.toJson(arraySetOf(currentGyroData.value))!!)
        }
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
            .addOnFailureListener { _ ->
                //TODO
                // If an error occurs, invoke the failure callback with the exception
            }

        val gpsData = gpsPrefs.getPref()
        if (gpsData != null) {
            val gpsSet: LinkedHashSet<Gps> = setAdapter?.fromJson(gpsData) as LinkedHashSet<Gps>
            gpsSet.add(currentGpsData.value)
            gpsPrefs.setPref(setAdapter.toJson(gpsSet).toString())
        } else {
            gpsPrefs.setPref(setAdapter?.toJson(arraySetOf(currentGpsData.value)).toString())
        }
    }

    private fun savePhotoDataLocally() {
        val photoData = photoPrefs.getPref()
        if (photoData != null) {
            val photoSet: LinkedHashSet<Photo> =
                setAdapter?.fromJson(photoData) as LinkedHashSet<Photo>
            photoSet.add(currentPhotoData.value)
            photoPrefs.setPref(setAdapter.toJson(photoSet).toString())
        } else {
            photoPrefs.setPref(setAdapter?.toJson(arraySetOf(currentPhotoData.value)).toString())
        }
    }

    private fun saveAndSendGyroData() {
        saveGyroDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            logsRepository.sendGyro(currentGyroData.value.toString())
            logs.add(GenericLog(LogType.GYRO, currentGyroData.value.copy()))
            Log.e("gyro","Gyro data saved and sent!")
        }
    }

    private fun saveAndSendGpsData() {
        if (hasNoLocationPermissions()) return

        saveGpsDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            val result = currentGpsData.value.toString()
            logsRepository.sendGps(result)
            logs.add(GenericLog(LogType.GPS, currentGpsData.value.copy()))
            Log.e("gps","GPS data saved and sent!")
        }
    }

    private fun saveAndSendPhotoData() {
        savePhotoDataLocally()
        CoroutineScope(Dispatchers.Main).launch {
            val result = currentPhotoData.value.toString()
            logsRepository.sendPhoto(result)
            Log.e("photo","Face detected, saved and sent!")
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

    internal fun processPicture(faceStatus: FaceStatus, bitmap: Bitmap?, timestamp: Long?) {
        Log.e("Face status", "This is ${faceStatus.name}")
        when (faceStatus) {
            FaceStatus.VALID -> {
                //Lets pretend I'm saving the Photo correctly here, shall we? ;)
                currentPhotoData.value.photo = bitmap.toString()
                currentPhotoData.value.timestamp = timestamp
                isFaceDetectedNow.value = true
            }

            else -> {
                isFaceDetectedNow.value = false
            }
        }
    }

}