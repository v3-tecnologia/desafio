import com.entregas.android_pleno_teste_v3.domain.data.FotoRequest
import com.entregas.android_pleno_teste_v3.domain.data.GiroscopioRequest
import com.entregas.android_pleno_teste_v3.domain.data.GpsRequest
import okhttp3.ResponseBody
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiService {

    @POST("/telemetry/gyroscope")
    fun enviarDadosGiroscopio(
        @Body dadosGiroscopio: GiroscopioRequest
    ): Result<Unit>

    @POST("/telemetry/gps")
    fun enviarDadosGps(
        @Body dadosGps: GpsRequest
    ): Result<Unit>

    @POST("/telemetry/photo")
    fun enviarDadosFoto(
        @Body dadosFoto: FotoRequest
    ): Result<Unit>
}
