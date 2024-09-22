package com.example.v3challenge.network

import com.squareup.moshi.Moshi
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory
import retrofit2.converter.moshi.MoshiConverterFactory

object ApiSettings {

    const val BASE_URL = "https://jsonplaceholder.typicode.com/"
    const val TEN_SECONDS = 2000L

    val moshi: Moshi =  Moshi.Builder().add(KotlinJsonAdapterFactory()).build()

    internal fun moshiFactory(): MoshiConverterFactory {
        return MoshiConverterFactory.create(moshi)
    }
}