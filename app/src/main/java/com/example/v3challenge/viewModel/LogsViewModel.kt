package com.example.v3challenge.viewModel

import android.app.Application
import androidx.lifecycle.ViewModel
import com.example.v3challenge.repository.LogsRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import javax.inject.Inject

@HiltViewModel
class LogsViewModel @Inject constructor(
    private val application: Application,
    private val logsRepository: LogsRepository): ViewModel() {

}