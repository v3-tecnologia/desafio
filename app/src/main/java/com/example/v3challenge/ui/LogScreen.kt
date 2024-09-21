package com.example.v3challenge.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.PowerSettingsNew
import androidx.compose.material3.FloatingActionButton
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.hilt.navigation.compose.hiltViewModel
import com.example.v3challenge.viewModel.LogsViewModel
import com.mutualmobile.composesensors.rememberGyroscopeSensorState

@Composable
fun LogScreen(viewModel: LogsViewModel = hiltViewModel()) {
    val log = viewModel.log
    val screenIsOn: MutableState<Boolean> = remember { mutableStateOf(false) }
    val scrollState = rememberScrollState()

    val sensorValue = rememberGyroscopeSensorState()
    viewModel.setGyroData(sensorValue)

    LaunchedEffect(key1 = true) {
        viewModel.startTimer()
        screenIsOn.value = true
    }

    @Composable
    fun screenContent() {
        Box(
            Modifier
                .fillMaxSize()
                .background(color = if (screenIsOn.value) Color.White else Color.Black)) {
            if (screenIsOn.value) {
                Column(modifier = Modifier.fillMaxSize().verticalScroll(scrollState)) {
                    Text(
                        text = log.value
                    )
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
    fun onOffButton(modifier: Modifier) {
        FloatingActionButton(modifier = modifier, onClick = { screenIsOn.value = !screenIsOn.value }) {
            if(screenIsOn.value) {
                Icon(Icons.Default.PowerSettingsNew, "")
            } else {
                Icon(Icons.Default.PowerSettingsNew, "")
            }

        }
    }

    //Main UI
    Box(modifier = Modifier.fillMaxSize()) {
        screenContent()
        onOffButton(
            Modifier
                .align(Alignment.BottomEnd)
                .padding(20.dp))
    }

}


