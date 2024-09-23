# Desafio Técnico V3

## ❤️ Bem vindos

Olá, tudo certo?

Seja bem vindo ao teste de seleção para novos desenvolvedores na V3!

Estamos honrados que você tenha chegado até aqui!

Prepare aquele ☕️ , e venha conosco codar e se divertir!

## Poxa, outro teste?

Nós sabemos que os processos de seleção podem ser ingratos! Você investe um tempão e no final pode não ser aprovado!

Aqui, nós presamos pela **transparência**!

Este teste tem um **propósito** bastante simples:

> Nós queremos avaliar como você consegue transformar problemas em soluções através de código!

**🚨 IMPORTANTE!** Se você entende que já possui algum projeto pessoal, ou contribuição em um projeto _open-source_ que contemple conhecimentos equivalentes aos que existem neste desafio, então, basta submeter o repositório explicando essa correlação!

## 🚀 Bora nessa!

Este é um teste para analisarmos como você desempenha ao entender, traduzir, resolver e entregar um código que resolve um problema.

### Dicas

- Documente seu projeto;
- Faça perguntas sobre os pontos que não ficaram claros para você;
- Mostre a sua linha de raciocínio;
- Trabalhe bem o seu README.md;
  - Explique até onde implementou;
  - Como o projeto pode ser executado;
  - Como pode-se testar o projeto;

### Como você deverá desenvolver?

1. Faça um _fork_ deste projeto em seu GitHub pessoal;
2. Realize as implementações de acordo com cada um dos níveis;
3. Faça pequenos _commits_;
4. Depois de sentir que fez o seu máximo, faça um PR para o repositório original.

🚨 **IMPORTANTE!** Não significa que você precisa implementar **todos os níveis** para ser aprovado no processo! Faça até onde se sentir confortável.

### Qual o tempo para entregar?

Quanto antes você enviar, mais cuidado podemos ter na revisão do seu teste. Mas sabemos que o dia a dia é corrido, faça de forma que fique confortável para você!

**Mas não desista! Envie até onde conseguir.**

## 💻 O Problema

Um dos nossos clientes ainda não consegue comprar o equipamento para colocar nos veículos de sua frota, mas ele quer muito utilizar a nossa solução.

Por isso, vamos fazer um MVP bastante simples para testar se, o celular do motorista poderia ser utilizado como o dispositivo de obtenção das informações.

> Parece fazer sentido certo? Ele possui vários mecanismos parecidos com o equipamento que oferecemos!

Sua missão ajudar na criação deste MVP para que possamos testar as frotas deste cliente.

Essa versão do produto será bastante simplificada. Queremos apenas criar as estruturas para obter algumas informações do seu dispositivo (Android) e armazená-la em um Banco de Dados.

Essas informações, depois de armazenadas devem estar disponíveis através de uma API para que este cliente integre com um Front-end já existente!

### Quais serão as informações que deverão ser coletadas?

1. **Dados de Giroscópio** - Estes dados devem retornar 3 valores (`x`, `y`, `z`). E devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foi coletado;
2. **Dados de GPS** - Estes dados devem retornar 2 valores (`latitude` , `longitude`). E também devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foram coletados;
3. **Uma foto** - Obter uma foto de uma das câmeras do dispositivo e enviá-la também junto com o `TIMESTAMP` em que foi coletada;

**🚨 É importante que se envie junto à essas informações um campo adicional, contendo uma identificação única do dispositivo, que pode ser seu endereço MAC.**

### Funcionamento

A aplicação Android deverá rodar em Background, e coletar e enviar as informações descritas a cada 10 segundos.

### Qual parte do desafio devo realizar?

Você deve realizar somente o desafio para a vaga que se candidatou.

Caso tenha sido a vaga de Android Embarcado, então resolva somente esta sessão.

Caso tenha sido a vaga de Backend, então resolva somente esta sessão.

---

# Desafio Android Embarcado

Você deverá criar uma aplicação que deverá coletar os dados e enviá-los para o servidor Back-end;

Lembre-se que essa é uma aplicação Android nativa, e não deve possuir qualquer tipo de interface com o usuário.

## Nível 1

Deve-se coletar os dados de acordo com as especificações, e armazená-los em um banco de dados local;

## Nível 2

Deve-se criar testes unitários para garantir o funcionamento das estruturas criadas;

## Nível 3

Deve-se enviar os dados obtidos a cada 10 segundos para uma API com a seguinte rota

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

## Nível 4

Deve-se realizar um _crop_ da foto obtida para que se consiga extrair somente um rosto. Caso a foto não tenha um rosto, ela não deverá ser enviada.

## Nível 5

Faça com que cada uma das requisições ocorra de forma paralela, e não de forma síncrona;

# Desafio Backend

Você deverá criar uma aplicação que irá receber os dados enviados pelo aplicativo.

Lembre-se essa aplicação precisa ser em GO!

## Nível 1

Deve-se criar uma API que receba requisições de acordo com os endpoints:

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

Deve-se garantir que os dados recebidos estão preenchidos corretamente.

Caso algum dado esteja faltando, então retorne uma mensagem de erro e um Status 400.

## Nível 2

Salve cada uma das informações em um banco de dados a sua escolha.

Salve estes dados de forma identificável e consistente;

## Nível 3

Crie testes unitários para cada arquivo da aplicação. Para cada nova implementação a seguir, também deve-se criar os testes.

## Nível 4

Crie um _container_ em _Docker_ que contenha a sua aplicação e o banco de dados utilizado nos testes.

## Nível 5

A cada foto recebida, deve-se utilizar o AWS Rekognition para comparar se a foto enviada é reconhecida com base nas fotos anteriores enviadas.

Se a foto enviada for reconhecida, retorne como resposta do `POST` um atributo que indique isso.

Utilize as fotos iniciais para realizar o treinamento da IA.
=======
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
9. [Licença](#licença)

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

## Licença

Este projeto é licenciado sob a Licença MIT. Veja o arquivo `LICENSE` para mais detalhes.