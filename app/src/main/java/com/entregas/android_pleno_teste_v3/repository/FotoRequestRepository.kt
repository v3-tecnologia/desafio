package com.entregas.android_pleno_teste_v3.repository

import ApiService
import com.entregas.android_pleno_teste_v3.domain.data.FotoRequest
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import retrofit2.HttpException
import java.io.IOException

class FotoRequestRepository(private val apiService: ApiService) {
    fun sendFotoRequest(
        fotoRequest: FotoRequest
    ): Flow<Result<Unit>> = flow {
        try {
            val response = apiService.enviarDadosFoto(fotoRequest)
            if (response.isSuccess) {
                emit(Result.success(Unit))
            } else {
                emit(Result.failure(Exception("Error: ${response} ")))
            }
        } catch (e: IOException) {
            emit(Result.failure(e))
        } catch (e: HttpException) {
            emit(Result.failure(e))
        }
    }
}
