package com.entregas.android_pleno_teste_v3.respositoryteste

import com.entregas.android_pleno_teste_v3.domain.PhotoRequestDataClass
import com.entregas.android_pleno_teste_v3.repository.FotoRequestRepository
import com.entregas.android_pleno_teste_v3.services.apiservice.ApiService
import com.entregas.android_pleno_teste_v3.services.apiservice.RetrofitClient
import io.mockk.coEvery
import io.mockk.every
import io.mockk.mockk
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test
import retrofit2.HttpException
import retrofit2.Response
import java.io.IOException

class FotoRequestRepositoryTest {
//Os testes não vai funcionar pois não temos uma BASE_URL para utilizar nos teste ,apenas os caminhos que foram colocados dentro da apiService
    private lateinit var fotoRequestRepository: FotoRequestRepository
    private val mockApiService: ApiService = mockk()
    private val mockRetrofitClient: RetrofitClient = mockk {
        every { apiService } returns mockApiService
    }

    @Before
    fun setUp() {
        // Setup repository with mocks
        fotoRequestRepository = FotoRequestRepository()
    }

    @Test
    fun `test sendPhotoRequest success`() = runTest {
        val fotoRequest = PhotoRequestDataClass("","")
        val response = Result.success(Unit) // Create a successful response
        coEvery { mockApiService.enviarFoto(fotoRequest) } returns response

        val result = fotoRequestRepository.sendPhotoRequest(fotoRequest).first()

        assertTrue(result.isSuccess)
        assertEquals(Unit, result.getOrNull())
    }

    @Test
    fun `test sendPhotoRequest HTTP exception`() = runTest {
        val fotoRequest = PhotoRequestDataClass("","")
        val httpException = HttpException(Response.error<Any>(500, okhttp3.ResponseBody.create(null, "Error")))
        coEvery { mockApiService.enviarFoto(fotoRequest) } throws httpException

        val result = fotoRequestRepository.sendPhotoRequest(fotoRequest).first()

        assertTrue(result.isFailure)
        assertTrue(result.exceptionOrNull() is HttpException)
    }

    @Test
    fun `test sendPhotoRequest IO exception`() = runTest {
        val fotoRequest = PhotoRequestDataClass("","")
        coEvery { mockApiService.enviarFoto(fotoRequest) } throws IOException()

        val result = fotoRequestRepository.sendPhotoRequest(fotoRequest).first()

        assertTrue(result.isFailure)
        assertTrue(result.exceptionOrNull() is IOException)
    }
}
