package com.entregas.android_pleno_teste_v3.application

import android.app.Application
import com.entregas.android_pleno_teste_v3.di.AppModule.appModule
import org.koin.android.ext.koin.androidContext
import org.koin.core.context.GlobalContext.startKoin

class MyApplication : Application() {
    override fun onCreate() {
        super.onCreate()
        startKoin {
            androidContext(this@MyApplication)
            modules(appModule)
        }
    }
}