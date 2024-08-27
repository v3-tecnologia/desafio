import androidx.room.Room
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.entregas.android_pleno_teste_v3.data.GyroscopeEntity
import com.entregas.android_pleno_teste_v3.data.LocationEntity
import com.entregas.android_pleno_teste_v3.data.PhotoEntity
import com.entregas.android_pleno_teste_v3.data.dao.GyroscopeDao
import com.entregas.android_pleno_teste_v3.data.dao.LocationDao
import com.entregas.android_pleno_teste_v3.data.dao.PhotoDao
import com.entregas.android_pleno_teste_v3.data.database.AppDatabase
import junit.framework.TestCase.assertEquals
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.runBlocking
import org.junit.After
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith

@ExperimentalCoroutinesApi
@RunWith(AndroidJUnit4::class)
class DataTestClass {

    private lateinit var gyroscopeDao: GyroscopeDao
    private lateinit var locationDao: LocationDao
    private lateinit var photoDao: PhotoDao
    private lateinit var database: AppDatabase

    @Before
    fun setUp() {
        database = Room.inMemoryDatabaseBuilder(
            ApplicationProvider.getApplicationContext(),
            AppDatabase::class.java
        ).build()
        locationDao = database.locationDao()
        gyroscopeDao = database.gyroscopeDao()
        photoDao = database.photoScopeDao()
    }

    @After
    fun finish() {
        database.close()
    }

    @Test
    fun testeSalvarLocationNoBancoLocalmente() = runBlocking {
        val locationEntity = LocationEntity(1, 0.0, 0.0, 200)
        locationDao.insert(locationEntity)

        val retrievedLocations = locationDao.getAllLocations()
        assertEquals(1, retrievedLocations.size)
        assertEquals(locationEntity, retrievedLocations.first())
    }

    @Test
    fun testeSalvarDadosGiroscopio() = runBlocking {
        val gyroscopeEntity = GyroscopeEntity(0, 0.0F, 0.0F, 0.0F, 1)
        gyroscopeDao.insert(gyroscopeEntity)

        val retrievedGyroscopes = gyroscopeDao.getAllGyroscopes()
        assertEquals(1, retrievedGyroscopes.size)
        assertEquals(gyroscopeEntity, retrievedGyroscopes.first())
    }

    @Test
    fun testeSalvarDadosDaFoto() = runBlocking {
        val photoEntity = PhotoEntity(2, "teste")
        photoDao.insert(photoEntity)

        val retrievedPhotos = photoDao.getAllPhotos()
        assertEquals(1, retrievedPhotos.size)
        assertEquals(photoEntity, retrievedPhotos.first())
    }
}
