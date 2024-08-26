package com.entregas.android_pleno_teste_v3.repository
import ApiService
import com.entregas.android_pleno_teste_v3.domain.data.GpsRequest
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import retrofit2.HttpException
import java.io.IOException

class GpsRequestRepository(private val apiService: ApiService) {

    fun sendGpsRequest(
        gpsRequest: GpsRequest
    ): Flow<Result<Unit>> = flow {
        try {
            val response = apiService.enviarDadosGps(gpsRequest)
            if (response.isSuccess) {
                emit(Result.success(Unit))
            } else {
                emit(Result.failure(Exception("Error: ${response}")))
            }
        } catch (e: IOException) {
            emit(Result.failure(e))
        } catch (e: HttpException) {
            emit(Result.failure(e))
        }
    }
}
