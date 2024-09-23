# Projeto V3 Teste

## Sumário
1. [Sobre o Projeto](#sobre-o-projeto)
2. [Estrutura do Projeto](#estrutura-do-projeto)
3. [Tecnologias Utilizadas](#tecnologias-utilizadas)
4. [Instalação e Configuração](#instalação-e-configuração)
5. [Execução do Projeto](#execução-do-projeto)
   - [Execução do Projeto Android](#execução-do-projeto-android)
   - [Execução do Projeto Backend](#execução-do-projeto-backend)
6. [Testes](#testes)
   - [Testes do Projeto Android](#testes-do-projeto-android)
   - [Testes do Projeto Backend](#testes-do-projeto-backend)
7. [Funcionalidades Implementadas](#funcionalidades-implementadas)
   - [Funcionalidades do Projeto Android](#funcionalidades-do-projeto-android)
   - [Funcionalidades do Projeto Backend](#funcionalidades-do-projeto-backend)
8. [Desafios e Soluções](#desafios-e-soluções)
9. [Como Contribuir](#como-contribuir)
10. [Licença](#licença)

## Sobre o Projeto

O **Projeto V3 Teste** é um MVP desenvolvido para demonstrar a coleta e processamento de dados de dispositivos Android, enviando essas informações para um backend construído em Go. O objetivo é substituir temporariamente o hardware dedicado com o celular do motorista para obter dados de giroscópio, GPS e capturar imagens. Esses dados são enviados ao backend para armazenamento e análise.

O projeto é dividido em duas partes principais:

- **AndroidEmbarcado**: Um aplicativo Android que coleta dados de sensores e envia essas informações para o backend.
- **BackendGo**: Uma API desenvolvida em Go para receber, armazenar e processar os dados enviados pelo aplicativo Android.

## Estrutura do Projeto

```go
v3_teste/
├── AndroidEmbarcado/
│   ├── build.gradle
│   ├── settings.gradle
│   ├── README.md
│   └── app/
│       ├── build.gradle
│       ├── src/
│       │   ├── main/
│       │   │   ├── AndroidManifest.xml
│       │   │   ├── java/
│       │   │   │   └── com/example/androidembarcado/
│       │   │   │       ├── MainActivity.kt
│       │   │   │       ├── database/
│       │   │   │       ├── model/
│       │   │   │       ├── network/
│       │   │   │       ├── repository/
│       │   │   │       ├── service/
│       │   │   │       ├── utils/
│       │   │   │       └── worker/
│       │   └── test/
│           └── java/
│               └── com/example/androidembarcado/
│                   ├── CameraServiceTest.kt
│                   ├── GpsServiceTest.kt
│                   ├── SensorServiceTest.kt
│                   └── TelemetryRepositoryTest.kt
├── BackendGo/
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   └── telemetry_controller.go
│   ├── models/
│   │   ├── gps.go
│   │   ├── gyroscope.go
│   │   └── photo.go
│   ├── repositories/
│   │   └── telemetry_repository.go
│   ├── routes/
│   │   └── routes.go
│   └── services/
│       ├── rekognition_service.go
│       └── telemetry_service.go
└── README.md    # Este README geral
```

## Tecnologias Utilizadas

### AndroidEmbarcado
- **Linguagem**: Kotlin
- **Bibliotecas**:
  - `Room Database` para armazenamento local.
  - `Retrofit` para chamadas de API.
  - `WorkManager` para tarefas em segundo plano.
  - `OpenCV` ou `ML Kit` para processamento de imagens.
  - `Coroutine` para operações assíncronas.

### BackendGo
- **Linguagem**: Go (Golang)
- **Bibliotecas**:
  - `Gorilla Mux` para roteamento de API.
  - `MongoDB Driver` para comunicação com o banco de dados.
  - `AWS SDK` para integração com AWS Rekognition.
- **Infraestrutura**:
  - Docker para containerização.
  - Docker Compose para gerenciamento dos serviços.

## Instalação e Configuração

1. **Clone o Repositório**
   ```bash
   git clone <link-do-repositorio>
   cd v3_teste
   ```

2. **Configuração do Projeto Android**
   - Abra o diretório `AndroidEmbarcado` no Android Studio.
   - Certifique-se de que as dependências do Gradle sejam sincronizadas.
   - Configure um dispositivo físico ou emulador Android com permissões necessárias.

3. **Configuração do Projeto Backend**
   - Certifique-se de ter o Docker instalado.
   - Navegue até o diretório `BackendGo` e configure as variáveis de ambiente necessárias:
     - `AWS_ACCESS_KEY`, `AWS_SECRET_KEY`, `DATABASE_URL`, `AWS_REGION`.
   - Execute o comando abaixo para iniciar a aplicação e o banco de dados:
     ```bash
     docker-compose up --build
     ```

## Execução do Projeto

### Execução do Projeto Android

1. No Android Studio, clique em **Run** para iniciar o aplicativo.
2. Certifique-se de que as permissões de GPS, câmera e armazenamento estejam habilitadas no dispositivo.
3. O aplicativo começará a coletar dados e enviá-los para o backend.

### Execução do Projeto Backend

1. Com o Docker em execução, a API estará disponível em `http://localhost:8080`.
2. Teste os endpoints utilizando o Postman ou outra ferramenta de sua escolha.

## Testes

### Testes do Projeto Android

1. Navegue até o diretório `AndroidEmbarcado` no Android Studio.
2. Execute os testes unitários para garantir a funcionalidade correta dos serviços e repositórios.
   - Os testes podem ser executados diretamente do `Android Studio` ou utilizando a linha de comando:
     ```bash
     ./gradlew test
     ```

### Testes do Projeto Backend

1. Navegue até o diretório `BackendGo`.
2. Execute os testes unitários:
   ```bash
   go test ./...
   ```

## Funcionalidades Implementadas

### Funcionalidades do Projeto Android

- **Coleta de Dados de Sensores**: Coleta dados de giroscópio e GPS.
- **Captura de Imagem**: Captura imagens utilizando a câmera do dispositivo.
- **Armazenamento Local**: Utiliza o Room Database para armazenamento local de dados.
- **Envio de Dados para o Backend**: Envia dados coletados para a API a cada 10 segundos.
- **Processamento de Imagem**: Realiza recorte (crop) de imagem para capturar apenas o rosto.

### Funcionalidades do Projeto Backend

- **Recepção de Dados**: Recebe e valida dados de giroscópio, GPS e imagens.
- **Armazenamento em Banco de Dados**: Armazena dados em MongoDB.
- **Integração com AWS Rekognition**: Compara imagens enviadas com imagens anteriores utilizando AWS Rekognition.
- **Containerização**: Projeto configurado para ser executado em containers Docker.

## Desafios e Soluções

1. **Coleta de Dados em Segundo Plano**: Foi utilizado `WorkManager` para garantir a coleta e envio de dados em segundo plano, mesmo quando o aplicativo não está em uso.
2. **Validação e Armazenamento**: Implementação de regras de validação no backend para garantir a integridade dos dados recebidos.
3. **Processamento de Imagens**: O uso de bibliotecas como OpenCV facilitou a manipulação e processamento de imagens no dispositivo Android.

## Como Contribuir

1. Faça um fork do projeto.
2. Crie uma nova branch (`git checkout -b feature/nova-feature`).
3. Faça suas alterações e commit (`git commit -am 'Adicionar nova feature'`).
4. Envie para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

## Licença

Este projeto é licenciado sob a Licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
