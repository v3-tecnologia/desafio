package com.entregas.android_pleno_teste_v3.datateste

import androidx.room.Room
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity
import com.entregas.android_pleno_teste_v3.data.LocationEntity
import com.entregas.android_pleno_teste_v3.data.dao.GyroscopeDao
import com.entregas.android_pleno_teste_v3.data.dao.LocationDao
import com.entregas.android_pleno_teste_v3.data.database.AppDatabase
import junit.framework.TestCase.assertNotNull
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.runBlocking
import org.junit.After
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith

@ExperimentalCoroutinesApi
class DataTestClass {
    private lateinit var gyroscopeDao : GyroscopeDao
    private lateinit var locationDao: LocationDao
    private lateinit var database: AppDatabase

    @Before
    fun setUp(){
        // Cria um banco de dados em memória
        database = Room.inMemoryDatabaseBuilder(
            ApplicationProvider.getApplicationContext(),
            AppDatabase::class.java
        ).build()
        locationDao = database.locationDao()
        gyroscopeDao = database.gyroscopeDao()
    }

    @After
    fun finish(){
        database.close()
    }

    @Test
     fun testeSalvarLocationNoBancoLocalmente() = runBlocking {
        val locationEntity = LocationEntity(1,0.0,0.0,200)
        locationDao.insert(locationEntity)

        // Verifica se o local foi salvo corretamente
        val retrievedLocation = locationDao.getAllLocations()
        assertNotNull(retrievedLocation)
    }
    @Test
    fun testeSalvarDadosGiroscopio() = runBlocking {
        val gyroscopeEntity = GyroscopeEntity(0, 0.0F,0.0F,0.0F,1)
        gyroscopeDao.insert(gyroscopeEntity)

        val getGyroscopeEntity= gyroscopeDao.getAllGyroscopes()
        assertNotNull(getGyroscopeEntity)
    }
}