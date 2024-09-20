package com.example.v3challenge.repository

import android.content.Context
import com.example.v3challenge.network.LogsInterface
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.android.scopes.ActivityScoped
import javax.inject.Inject

@ActivityScoped
class LogsRepository @Inject constructor(
    private val logRepository: LogsInterface,
    @ApplicationContext val context: Context) {

}