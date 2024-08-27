package com.entregas.android_pleno_teste_v3.data.database

import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase
import com.entregas.android_pleno_teste_v3.data.dao.GyroscopeDao
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity
import com.entregas.android_pleno_teste_v3.data.dao.LocationDao
import com.entregas.android_pleno_teste_v3.data.LocationEntity
import com.entregas.android_pleno_teste_v3.data.dao.PhotoDao

@Database(entities = [LocationEntity::class, GyroscopeEntity::class], version = 1)
abstract class AppDatabase : RoomDatabase() {
    abstract fun locationDao(): LocationDao
    abstract fun gyroscopeDao(): GyroscopeDao
    abstract fun photoScopeDao() : PhotoDao
            companion object {
            @Volatile
            private var INSTANCE: AppDatabase? = null

            fun getDatabase(context: Context): AppDatabase {
                return INSTANCE ?: synchronized(this) {
                    val instance = Room.databaseBuilder(
                        context.applicationContext,
                        AppDatabase::class.java,
                        "app_database"
                    ).build()
                    INSTANCE = instance
                    instance
                }
            }
    }
}
