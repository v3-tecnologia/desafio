Descrição
Este projeto é uma aplicação Android que opera em segundo plano para coletar dados de localização e giroscópio. Os dados são armazenados localmente usando o Room e enviados para um backend via Retrofit. O aplicativo também possui a funcionalidade de capturar fotos e verificar se essas fotos contêm rostos.

Funcionalidades
Serviço em Segundo Plano: Coleta dados de localização e giroscópio periodicamente.
Persistência de Dados: Armazena dados de localização e giroscópio localmente usando o Room.
Envio de Dados: Utiliza Retrofit para enviar dados, incluindo o endereço MAC do dispositivo, para um backend.
Captura de Fotos: Tira fotos e verifica se elas contêm rostos.
Notificações: Notifica o usuário que o serviço está em execução.
Requisitos
Android Studio
Gradle
Permissões de localização e acesso ao serviço em primeiro plano
Configuração do Projeto
Clonagem do Repositório Clone o repositório usando o comando:

bash
Copy code
git clone https://github.com/SEU_USUARIO/NOVO_REPOSITORIO.git
Configuração das Permissões Certifique-se de que o seu arquivo AndroidManifest.xml inclui as permissões necessárias:

xml
<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
<uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION" />
<uses-permission android:name="android.permission.INTERNET" />
<uses-permission android:name="android.permission.FOREGROUND_SERVICE" />
<uses-permission android:name="android.permission.RECEIVE_BOOT_COMPLETED" />
Configuração do Serviço O BackgroundService coleta dados de localização e giroscópio e envia para o backend. O serviço deve ser iniciado a partir da MainActivity se as permissões de localização forem concedidas.

Configuração do Retrofit Configure o Retrofit para enviar dados para o backend.

Configuração do Room Configure o Room para armazenar dados localmente. Defina as entidades e o banco de dados conforme necessário.

Autenticação e Autorização Se houver problemas com autenticação do GitHub, verifique suas configurações de chave SSH ou tokens de acesso pessoal.

Execução do Projeto
Compile o Projeto: No Android Studio, selecione Build > Rebuild Project.
Execute o Aplicativo: Inicie o aplicativo no seu dispositivo ou emulador.
Testes
Testes Unitários: Foram realizados testes unitários para garantir que as funcionalidades individuais do aplicativo estão funcionando corretamente. Esses testes verificam o comportamento das classes e métodos isoladamente, garantindo que cada unidade de código execute a tarefa esperada.

Testes de Integração: Testes de integração foram executados para assegurar que os diferentes componentes do aplicativo funcionam bem juntos. Isso inclui a integração do serviço em segundo plano com o banco de dados Room e a comunicação com o backend via Retrofit.

Testes de API: Foram realizados testes de API para garantir que a comunicação com o backend esteja correta. Estes testes verificam se as requisições e respostas da API estão sendo processadas como esperado, e se os dados são enviados e recebidos corretamente.

Contribuições
Sinta-se à vontade para enviar pull requests e contribuir para o projeto.

