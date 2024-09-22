package com.example.v3challenge.di

import android.content.Context
import com.example.v3challenge.network.ApiSettings
import com.example.v3challenge.network.ApiSettings.moshiFactory
import com.example.v3challenge.network.LogsInterface
import com.example.v3challenge.repository.LogsRepository
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.components.SingletonComponent
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
class LogModule {

    @Singleton
    @Provides
    fun provideLogRepository(
        logs: LogsInterface,
        @ApplicationContext context: Context
    ) = LogsRepository(logs, context)

    @Singleton
    @Provides
    fun logService(): LogsInterface {
        return Retrofit.Builder()
            .baseUrl(ApiSettings.BASE_URL)
            .addConverterFactory(moshiFactory())
            .client(OkHttpClient.Builder().build())
            .build()
            .create(LogsInterface::class.java)
    }
}