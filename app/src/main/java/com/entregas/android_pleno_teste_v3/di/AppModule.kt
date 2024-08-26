package com.entregas.android_pleno_teste_v3.di

import com.entregas.android_pleno_teste_v3.repository.FotoRequestRepository
import com.entregas.android_pleno_teste_v3.repository.GiroscopioRequestRepository
import com.entregas.android_pleno_teste_v3.services.apiservice.ApiService
import com.entregas.android_pleno_teste_v3.utils.Constants.Companion.BASE_URL
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import org.koin.dsl.module

object AppModule {
    val appModule = module {
        single {
            OkHttpClient.Builder().build()
        }

        single {
            Retrofit.Builder()
                .baseUrl(BASE_URL)
                .client(get())
                .addConverterFactory(GsonConverterFactory.create())
                .build()
        }

        single {
            get<Retrofit>().create(ApiService::class.java)
        }
        single {
            FotoRequestRepository(get())
        }
        single{
            GiroscopioRequestRepository(get())
        }

        single{
            GiroscopioRequestRepository(get())
        }
    }
}
