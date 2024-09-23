package com.example.androidembarcado.database

import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase
import com.example.androidembarcado.model.GpsData
import com.example.androidembarcado.model.GyroscopeData
import com.example.androidembarcado.model.PhotoData

@Database(entities = [GyroscopeData::class, GpsData::class, PhotoData::class], version = 1)
abstract class AppDatabase : RoomDatabase() {

    abstract fun gyroscopeDao(): GyroscopeDao
    abstract fun gpsDao(): GpsDao
    abstract fun photoDao(): PhotoDao

    companion object {
        @Volatile private var INSTANCE: AppDatabase? = null

        fun getInstance(context: Context): AppDatabase {
            return INSTANCE ?: synchronized(this) {
                val instance = Room.databaseBuilder(
                    context.applicationContext,
                    AppDatabase::class.java,
                    "telemetry_database"
                ).build()
                INSTANCE = instance
                instance
            }
        }
    }
}
