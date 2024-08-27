## Funcionalidades

- **Serviço em Segundo Plano**: Coleta dados de localização e giroscópio periodicamente a cada 10 segundos.
- **Persistência de Dados**: Armazena dados de localização e giroscópio localmente usando Room.
- **Envio de Dados**: Envia dados para um backend via Retrofit, incluindo o ID único do dispositivo.
- **Captura de Fotos**: Tira fotos e verifica se elas contêm rostos.
- **Notificações**: Informa o usuário quando o serviço está em execução.

## Requisitos

- **Android Studio**: Ferramenta de desenvolvimento para construir e testar o aplicativo.
- **Gradle**: Sistema de automação de build usado para gerenciar dependências e compilar o projeto.
- **Permissões**: O aplicativo requer permissões para acessar dados de localização, arquivos da galeria e para execução de serviços em primeiro plano.

## Configuração do Projeto

### Clonagem do Repositório

Clone o repositório do projeto para o seu ambiente de desenvolvimento.

### Configuração das Permissões

Certifique-se de incluir as seguintes permissões no arquivo `AndroidManifest.xml` do projeto:

- **Localização**: Permissões para acessar dados de localização.
- **Serviço em Primeiro Plano**: Permissões para garantir que o serviço de coleta de dados continue funcionando em segundo plano.
- **Acesso a Arquivos da Galeria**: Permissões para ler e escrever arquivos na galeria, conforme a versão do Android.

### Configuração dos Serviços

- **Serviço de Localização e Giroscópio**: Um serviço dedicado coleta dados de localização e giroscópio periodicamente, a cada 10 segundos. Este serviço deve ser iniciado a partir da `MainActivity`, garantindo que as permissões necessárias sejam concedidas.

- **Serviço de Captura de Fotos**: Um segundo serviço é responsável por capturar fotos e verificar se essas fotos contêm rostos. Esse serviço opera separadamente do serviço de coleta de dados de localização e giroscópio.

### Configuração do Room

Room é utilizado para armazenar dados de localização e giroscópio localmente a cada 10 segundos. Garanta que o Room esteja configurado para persistir esses dados, permitindo o acesso posterior conforme necessário.

### Configuração do Retrofit

Retrofit é utilizado para enviar dados para um backend, incluindo o ID único do dispositivo, que serve como uma identificação única. Este ID muda somente se o dispositivo for restaurado ou sofrer outras alterações significativas.

## Notas Adicionais

- **Permissões**: O aplicativo verifica a versão do Android e solicita permissões apropriadas para acessar arquivos da galeria, conforme a versão do sistema operacional.
- **Serviços**: O aplicativo utiliza dois serviços distintos: um para coleta de dados de localização e giroscópio, e outro para captura de fotos.
