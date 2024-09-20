package com.example.v3challenge.ui

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.example.v3challenge.viewModel.LogsViewModel
import androidx.hilt.navigation.compose.hiltViewModel

@Composable
fun LogScreen(viewModel: LogsViewModel = hiltViewModel(),) {
    Text(
        text = "Hello Android"
    )
}