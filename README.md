Android Pleno Teste V3
Descrição
Este projeto é uma aplicação Android que inclui um serviço em segundo plano para coleta de dados de localização e giroscópio. Os dados são armazenados usando o Room e enviados para um backend via Retrofit.

Funcionalidades
Serviço em Segundo Plano: Coleta dados de localização e giroscópio periodicamente.
Persistência de Dados: Armazena dados de localização e giroscópio usando o Room.
Envio de Dados: Utiliza Retrofit para enviar dados para um backend.
Notificações: Notifica o usuário que o serviço está em execução.
Requisitos
Android Studio
Gradle
Permissões de localização e acesso ao serviço em primeiro plano
Configuração do Projeto
Passo 1: Clonagem do Repositório
Clone o repositório usando o comando:

git clone https://github.com/SEU_USUARIO/NOVO_REPOSITORIO.git

Passo 2: Configuração das Dependências
O projeto usa o Room e Retrofit para persistência de dados e comunicação com o backend. Certifique-se de que as seguintes dependências estão incluídas no seu arquivo build.gradle:

// Room

implementation "androidx.room:room-runtime:2.4.3"
annotationProcessor "androidx.room:room-compiler:2.4.3"

// Retrofit

implementation "com.squareup.retrofit2:retrofit:2.9.0"
implementation "com.squareup.retrofit2:converter-gson:2.9.0"

Passo 3: Configuração das Permissões
Certifique-se de que o seu arquivo AndroidManifest.xml inclui as permissões necessárias:

<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
<uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION" />
<uses-permission android:name="android.permission.INTERNET" />
<uses-permission android:name="android.permission.FOREGROUND_SERVICE" />
<uses-permission android:name="android.permission.RECEIVE_BOOT_COMPLETED" />
Passo 4: Configuração do Serviço
O BackgroundService coleta dados de localização e giroscópio e envia para o backend. O serviço deve ser iniciado a partir da MainActivity se as permissões de localização forem concedidas.

Passo 5: Configuração do Retrofit
Configure o Retrofit para enviar dados para o backend. Adicione as configurações do Retrofit em uma classe de configuração:

object RetrofitClient {
    private const val BASE_URL = "https://example.com/api/"

    val instance: Retrofit by lazy {
        Retrofit.Builder()
            .baseUrl(BASE_URL)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }
}
Passo 6: Configuração do Room
Configure o Room para armazenar dados localmente. Defina as entidades e o banco de dados no seu código.

Passo 7: Autenticação e Autorização
Se você estiver enfrentando problemas com autenticação do GitHub, verifique suas configurações de chave SSH ou tokens de acesso pessoal.

Execução do Projeto
Compile o projeto: No Android Studio, selecione Build > Rebuild Project.
Execute o aplicativo: Inicie o aplicativo no seu dispositivo ou emulador.
Testes
Adicione testes unitários e de integração para garantir que o serviço e a comunicação com o backend estão funcionando corretamente. Utilize o framework de testes do Android e ferramentas como Espresso para testes de UI.

Contribuições
Sinta-se à vontade para enviar pull requests e contribuir para o projeto. Para mais informações, consulte o arquivo CONTRIBUTING.md.

Licença
Este projeto está licenciado sob a Licença MIT.
