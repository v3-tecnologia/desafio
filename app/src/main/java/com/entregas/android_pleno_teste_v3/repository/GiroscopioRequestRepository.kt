package com.entregas.android_pleno_teste_v3.repository

import ApiService
import com.entregas.android_pleno_teste_v3.domain.data.GiroscopioRequest
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import retrofit2.HttpException
import java.io.IOException

class GyroscopeRequestRepository(private val apiService: ApiService) {

    fun sendGyroscopeRequest(
        gyroscopeRequest: GiroscopioRequest
    ): Flow<Result<Unit>> = flow {
        try {
            val response = apiService.enviarDadosGiroscopio(gyroscopeRequest)
            if (response.isSuccess) {
                emit(Result.success(Unit))
            } else {
                emit(Result.failure(Exception("Error: ${response} ${response}")))
            }
        } catch (e: IOException) {
            emit(Result.failure(e))
        } catch (e: HttpException) {
            emit(Result.failure(e))
        }
    }
}
