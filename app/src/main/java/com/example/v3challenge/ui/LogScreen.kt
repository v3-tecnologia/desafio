package com.example.v3challenge.ui

import androidx.camera.core.CameraSelector
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.DeleteSweep
import androidx.compose.material.icons.filled.FlipCameraIos
import androidx.compose.material.icons.filled.PowerSettingsNew
import androidx.compose.material3.Button
import androidx.compose.material3.FloatingActionButton
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.LocalLifecycleOwner
import com.example.v3challenge.utils.CameraManager
import com.example.v3challenge.viewModel.LogsViewModel
import com.mutualmobile.composesensors.rememberGyroscopeSensorState

@Composable
fun LogScreen(viewModel: LogsViewModel = hiltViewModel()) {
    val context = LocalContext.current
    val lifecycleOwner = LocalLifecycleOwner.current
    val logs = viewModel.logs
    val screenIsOn: MutableState<Boolean> = remember { mutableStateOf(false) }
    val scrollState = rememberLazyListState()
    val sensorValue = rememberGyroscopeSensorState()
    val cameraSelectorOption: MutableState<Int> =
        remember { mutableIntStateOf(CameraSelector.LENS_FACING_FRONT) }

    val cameraManager = CameraManager(
        context,
        lifecycleOwner,
        cameraSelectorOption,
        viewModel::processPicture
    )

    viewModel.setGyroData(sensorValue)

    LaunchedEffect(key1 = true) {
        cameraManager.startCamera()
        viewModel.startTimer()
        screenIsOn.value = true
    }

    LaunchedEffect(key1 = logs.size) {
        //Scrolls to bottom everytime the log changes
        if (logs.isNotEmpty()) {
            scrollState.scrollToItem(logs.size - 1)
        }
    }

    RequestPermissions({}, {}, {})

    @Composable
    fun screenContent() {
        Box(
            Modifier
                .fillMaxSize()
                .background(color = if (screenIsOn.value) Color.White else Color.Black)
        ) {
            if (screenIsOn.value) {
                LazyColumn(
                    modifier = Modifier
                        .fillMaxSize()
                        .padding(bottom = 100.dp),
                    state = scrollState
                ) {
                    logs.forEach { log ->
                        item {
                            Text(
                                text = log,
                                color = Color.Black,
                                fontSize = 14.sp
                            )
                        }
                    }
                }
            } else {
                Text(
                    modifier = Modifier.align(alignment = Alignment.Center),
                    text = "Screen OFF",
                    color = Color.White,
                    fontSize = 20.sp
                )
            }
        }
    }

    @Composable
    fun OnOffButton(modifier: Modifier) {
        FloatingActionButton(
            modifier = modifier,
            onClick = { screenIsOn.value = !screenIsOn.value }) {
            Icon(Icons.Default.PowerSettingsNew, "")
        }
    }

    @Composable
    fun CameraSwitchButton() {
        Button(
            onClick = {
                cameraManager.changeCameraSelector()
            }) {
            Icon(Icons.Default.FlipCameraIos, "")
        }
    }

    @Composable
    fun ClearLogButton(modifier: Modifier) {
        Button(
            modifier = modifier,
            onClick = {
                viewModel.logs.clear()
            }) {
            Icon(Icons.Default.DeleteSweep, "")
        }

    }

    //Main UI
    Box(modifier = Modifier.fillMaxSize()) {
        screenContent()
        OnOffButton(
            Modifier
                .align(Alignment.BottomEnd)
                .padding(20.dp)
        )
        if (screenIsOn.value) {
            val cameraUsed = if (cameraSelectorOption.value == 0) "FRONT" else "REAR"
            Column(
                modifier = Modifier
                    .width(100.dp)
                    .align(Alignment.BottomStart)
                    .padding(bottom = 20.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Text(
                    text = "$cameraUsed camera",
                    color = Color.Blue,
                    fontSize = 14.sp
                )
                CameraSwitchButton()
            }
            ClearLogButton(
                Modifier
                    .align(Alignment.BottomCenter)
                    .padding(bottom = 20.dp)
            )
        }
    }

}


