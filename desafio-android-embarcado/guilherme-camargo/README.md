# Desafio Android Embarcado

## Índice

- [Desafio Android Embarcado](#desafio-android-embarcado)
  - [Índice](#índice)
  - [O Problema](#o-problema)
    - [Objetivo do MVP](#objetivo-do-mvp)
    - [Informações a Serem Coletadas](#informações-a-serem-coletadas)
  - [Níveis de Desenvolvimento](#níveis-de-desenvolvimento)
    - [Nível 1: Coleta e Armazenamento de Dados](#nível-1-coleta-e-armazenamento-de-dados)
    - [Nível 2: Testes Unitários](#nível-2-testes-unitários)
    - [Nível 3: Envio dos Dados para a API](#nível-3-envio-dos-dados-para-a-api)
    - [Nível 4: Processamento de Imagem](#nível-4-processamento-de-imagem)
    - [Nível 5: Requisições Paralelas](#nível-5-requisições-paralelas)
    - [AppUpdateReceiver](#appupdatereceiver)
    - [BackgroundService](#backgroundservice)
    - [BootReceiver](#bootreceiver)
    - [CustomCameraManager](#customcameramanager)
    - [GyroscopeManager](#gyroscopemanager)
    - [LocationManager](#locationmanager)
    - [DatabaseManager](#databasemanager)
    - [LogWriter](#logwriter)
  - [Como Iniciar o Aplicativo Usando ADB](#como-iniciar-o-aplicativo-usando-adb)
    - [Iniciando o Serviço em Primeiro Plano](#iniciando-o-serviço-em-primeiro-plano)
  - [Configurando Variáveis de Ambiente para Android SDK e ADB](#configurando-variáveis-de-ambiente-para-android-sdk-e-adb)
    - [Passos para Configuração](#passos-para-configuração)
    - [1. Localize o Caminho do Android SDK](#1-localize-o-caminho-do-android-sdk)
    - [2. Configurando no Windows](#2-configurando-no-windows)
    - [3. Configurando no macOS e Linux](#3-configurando-no-macos-e-linux)
    - [4. Verificando a Configuração](#4-verificando-a-configuração)

## O Problema

Um dos nossos clientes ainda não consegue comprar o equipamento para colocar nos veículos de sua frota, mas ele quer muito utilizar a nossa solução.

Para resolver essa questão, estamos criando um MVP simples para testar se o celular do motorista pode ser utilizado como o dispositivo de obtenção das informações. A ideia é que o celular, que possui mecanismos semelhantes ao equipamento que oferecemos, possa ser utilizado para coletar dados relevantes enquanto o motorista está em operação.

### Objetivo do MVP

Este MVP visa coletar dados críticos do dispositivo Android do motorista, como giroscópio, GPS e uma foto, e armazená-los em um banco de dados local. Essas informações, posteriormente, serão disponibilizadas através de uma API para integração com um front-end existente do cliente.

### Informações a Serem Coletadas

1. **Dados de Giroscópio**: Três valores (x, y, z) juntamente com o TIMESTAMP do momento da coleta.
2. **Dados de GPS**: Dois valores (latitude, longitude) juntamente com o TIMESTAMP do momento da coleta.
3. **Uma Foto**: Uma foto capturada de uma das câmeras do dispositivo, juntamente com o TIMESTAMP da captura.
4. **Identificação do Dispositivo**: Um campo adicional com uma identificação única do dispositivo, como o endereço MAC.

## Níveis de Desenvolvimento

A aplicação foi desenvolvida em níveis, cada um com seus próprios objetivos e requisitos. Abaixo estão os níveis com uma checklist para o que foi feito e o que ainda precisa ser feito.

### Nível 1: Coleta e Armazenamento de Dados

- [x] Coletar dados de giroscópio (x, y, z) e armazená-los com TIMESTAMP.
- [x] Coletar dados de GPS (latitude, longitude) e armazená-los com TIMESTAMP.
- [x] Capturar uma foto e armazená-la com TIMESTAMP.
- [x] Armazenar a identificação única do dispositivo (endereço MAC ou Device ID).
- [x] Armazenar todas as informações coletadas em um banco de dados local.

### Nível 2: Testes Unitários

- [Pendente] Criar testes unitários para garantir o funcionamento das estruturas de coleta de dados.
- [Pendente] Criar testes unitários para o armazenamento no banco de dados local.

### Nível 3: Envio dos Dados para a API

- [x] Enviar os dados de giroscópio para a API via POST /telemetry/gyroscope a cada 10 segundos.
- [x] Enviar os dados de GPS para a API via POST /telemetry/gps a cada 10 segundos.
- [x] Enviar a foto para a API via POST /telemetry/photo a cada 10 segundos.

### Nível 4: Processamento de Imagem

- [ ] Realizar um crop da foto para extrair somente o rosto.
- [ ] Verificar se a foto contém um rosto; caso contrário, não enviar a foto.

### Nível 5: Requisições Paralelas

- [x] Implementar a execução paralela das requisições para giroscópio, GPS e foto.
- [x] Garantir que as requisições não sejam feitas de forma síncrona.

### AppUpdateReceiver

**Descrição**:
O `AppUpdateReceiver` é um `BroadcastReceiver` que escuta as atualizações do aplicativo. Ele é responsável por reiniciar o `BackgroundService` sempre que o aplicativo é atualizado, garantindo que o serviço continue rodando após uma atualização.

**Função Principal**:

- Receber uma transmissão quando o pacote do aplicativo for substituído ou atualizado.
- Reiniciar o `BackgroundService` após a atualização do aplicativo.

**Métodos Principais**:

- `onReceive(context: Context, intent: Intent)`: Método chamado quando o `BroadcastReceiver` recebe uma intenção (intent). Verifica se o evento de atualização do pacote foi recebido e reinicia o serviço de background.

### BackgroundService

**Descrição**:
O `BackgroundService` é um serviço Android que roda em segundo plano e é responsável por coletar dados do giroscópio, localização GPS, e capturar imagens. Esses dados são periodicamente armazenados em um banco de dados local e podem ser utilizados para monitoramento ou análises posteriores. O serviço é projetado para rodar continuamente, mesmo após reinicializações do dispositivo.

**Funcionalidades Principais**:

- **Coleta de Dados do Giroscópio**: Coleta e armazena os valores das coordenadas x, y, z juntamente com o timestamp de coleta.
- **Coleta de Dados de Localização**: Coleta e armazena a latitude e longitude juntamente com o timestamp de coleta.
- **Captura de Imagens**: Captura uma imagem usando a câmera do dispositivo, a converte para Base64, e a armazena juntamente com o timestamp de captura.
- **Identificação Única do Dispositivo**: Captura o `ANDROID_ID` do dispositivo para ser utilizado como um identificador único, garantindo que os dados coletados sejam associados ao dispositivo específico.
- **Serviço em Primeiro Plano**: Pode ser executado em primeiro plano para evitar que o sistema encerre o serviço.

**Métodos Principais**:

- `onCreate()`: Inicializa todos os gerenciadores necessários (`GyroscopeManager`, `LocationManager`, `CustomCameraManager`) e o banco de dados, e captura a identificação única do dispositivo. Inicia a coleta de dados em intervalos regulares.
- `onStartCommand(intent: Intent?, flags: Int, startId: Int)`: Garante que o serviço continue rodando após ser iniciado.
- `onDestroy()`: Limpa os recursos e para a coleta de dados ao encerrar o serviço.
- `executeRoutine()`: Método que executa a rotina de coleta de dados (giroscópio, localização, e câmera) a cada 10 segundos.
- `getDeviceId()`: Método que captura a identificação única do dispositivo utilizando o `ANDROID_ID`.

**Tratamento de Exceções**:

- O serviço possui tratamento de exceções robusto para capturar e registrar erros durante a inicialização dos componentes e a execução da rotina de coleta de dados. Os logs de erro são gerenciados pela classe `LogWriter`.

### BootReceiver

**Descrição**:
O `BootReceiver` é um `BroadcastReceiver` que inicia o `BackgroundService` automaticamente após o dispositivo ser ligado. Isso garante que o serviço de background comece a rodar sempre que o dispositivo for reiniciado.

**Função Principal**:

- Iniciar o `BackgroundService` quando o dispositivo termina de iniciar.

**Métodos Principais**:

- `onReceive(context: Context, intent: Intent)`: Método chamado quando o `BroadcastReceiver` recebe uma intenção (intent). Verifica se o evento de boot foi recebido e reinicia o serviço de background.

### CustomCameraManager

**Descrição**:
O `CustomCameraManager` gerencia as operações da câmera, incluindo a captura de imagens em segundo plano. Ele é usado pelo `BackgroundService` para capturar imagens sem a necessidade de interação do usuário ou uma pré-visualização.

**Função Principal**:

- Abrir a câmera, capturar imagens, e retornar os dados da imagem em formato Base64.
- Fechar a câmera corretamente após a captura.

**Métodos Principais**:

- `initializeCameraAndTakePicture()`: Inicializa a câmera, captura uma imagem, e retorna a imagem capturada em formato Base64.
- `captureImage()`: Configura a captura da imagem e chama o método que converte a imagem para Base64.
- `saveImageAsBase64()`: Converte a imagem capturada em uma string Base64 para fácil manipulação.
- `stopCamera()`: Fecha a sessão da câmera e libera os recursos.

### GyroscopeManager

**Descrição**:
O `GyroscopeManager` é responsável por coletar dados do giroscópio do dispositivo. Ele registra um listener que recebe as leituras do giroscópio e armazena os valores das coordenadas x, y e z.

**Função Principal**:

- Monitorar e capturar dados do giroscópio em tempo real.

**Métodos Principais**:

- `onSensorChanged(event: SensorEvent?)`: Método chamado sempre que há uma nova leitura do giroscópio. Atualiza os valores x, y e z.
- `unregisterListener()`: Desregistra o listener do giroscópio para parar de receber atualizações.

### LocationManager

**Descrição**:
O `LocationManager` gerencia a coleta de dados de localização, utilizando o serviço de localização do Google para capturar a latitude e a longitude do dispositivo.

**Função Principal**:

- Coletar periodicamente a localização do dispositivo (latitude e longitude).

**Métodos Principais**:

- `startLocationUpdates()`: Inicia a captura contínua da localização do dispositivo.
- `stopLocationUpdates()`: Para a captura de localização e libera os recursos.
- `onLocationResult(locationResult: LocationResult)`: Callback que é chamado quando uma nova localização é recebida.

### DatabaseManager

**Descrição**:
O `DatabaseManager` é uma classe que gerencia o banco de dados local do aplicativo. Ele é responsável por criar e manter três tabelas que armazenam os dados coletados pelo `BackgroundService`: dados do giroscópio, dados de localização e dados de imagem. Essa classe também lida com a inserção dos dados nas tabelas e a criação do banco de dados quando o aplicativo é iniciado pela primeira vez.

**Tabelas Gerenciadas**:

- **Tabela `GyroscopeData`**: Armazena os valores x, y, z do giroscópio juntamente com o timestamp da coleta.
- **Tabela `LocationData`**: Armazena a latitude e longitude do dispositivo juntamente com o timestamp da coleta.
- **Tabela `ImageData`**: Armazena a imagem capturada, em formato Base64, juntamente com o timestamp da captura.

**Métodos Principais**:

- `onCreate(db: SQLiteDatabase?)`: Cria as três tabelas (`GyroscopeData`, `LocationData`, `ImageData`) quando o banco de dados é criado pela primeira vez.
- `onUpgrade(db: SQLiteDatabase?, oldVersion: Int, newVersion: Int)`: Atualiza o banco de dados, caso haja uma nova versão, recriando as tabelas se necessário.
- `insertGyroscopeData(xValue: Float, yValue: Float, zValue: Float)`: Insere dados do giroscópio na tabela `GyroscopeData`.
- `insertLocationData(latitude: Double, longitude: Double)`: Insere dados de localização na tabela `LocationData`.
- `insertImageData(imageBase64: String)`: Insere dados de imagem na tabela `ImageData`.

**Tratamento de Exceções**:

- Todos os métodos de inserção e criação de tabelas possuem tratamento de exceções para garantir que erros sejam capturados e registrados. O `LogWriter` é utilizado para armazenar logs de sucesso e falha nas operações do banco de dados.

**Uso do `LogWriter`**:

- Todos os erros e operações bem-sucedidas (como a criação de tabelas e a inserção de dados) são registrados usando o `LogWriter`, proporcionando uma trilha de auditoria completa para facilitar o diagnóstico de problemas.

### LogWriter

**Descrição**:
O `LogWriter` é uma classe utilitária para gravar logs em um arquivo `.txt` no dispositivo. Ele é utilizado por outras classes do projeto para registrar eventos importantes e mensagens de erro.

**Função Principal**:

- Escrever mensagens de log em um arquivo `.txt` no dispositivo, com suporte a mensagens de erro e logs informativos.

**Métodos Principais**:

- `writeLog(tag: String, message: String)`: Escreve uma mensagem de log em um arquivo. Se o arquivo já existir, ele adiciona a mensagem ao final do arquivo.
- `generateLogFileName()`: Gera um nome de arquivo com base na data e hora atuais para organizar os logs.

## Como Iniciar o Aplicativo Usando ADB

O `BackgroundService` do aplicativo pode ser iniciado diretamente a partir do shell ADB, o que é útil para testes ou para iniciar o serviço em dispositivos onde a interface gráfica não está disponível.

### Iniciando o Serviço em Primeiro Plano

Para iniciar o `BackgroundService` em primeiro plano, use o seguinte comando ADB:

```bash
adb shell am start-foreground-service -n br.com.camargoguilherme.desafio_android_embarcado/.BackgroundService
```

## Configurando Variáveis de Ambiente para Android SDK e ADB

Se você ainda não configurou as variáveis de ambiente para o Android SDK e ADB, siga os passos abaixo para garantir que possa utilizar o comando ADB e outras ferramentas do Android SDK diretamente no terminal.

### Passos para Configuração

### 1. Localize o Caminho do Android SDK

O Android SDK geralmente é instalado em uma das seguintes localizações:

- **Windows**: `C:\Users\<SeuUsuário>\AppData\Local\Android\Sdk`
- **macOS**: `/Users/<SeuUsuário>/Library/Android/sdk`
- **Linux**: `/home/<SeuUsuário>/Android/Sdk`

### 2. Configurando no Windows

1. **Abra as Configurações do Sistema**:

   - Vá em `Painel de Controle` > `Sistema e Segurança` > `Sistema` > `Configurações Avançadas do Sistema`.
   - Clique em `Variáveis de Ambiente`.

2. **Adicione o Caminho do SDK**:

   - Na seção `Variáveis do sistema`, selecione `Path` e clique em `Editar`.
   - Clique em `Novo` e adicione o caminho para o diretório `platform-tools` do Android SDK. Exemplo:
     ```
     C:\Users\<SeuUsuário>\AppData\Local\Android\Sdk\platform-tools
     ```

3. **Salve as Alterações** e feche as janelas.

### 3. Configurando no macOS e Linux

1. **Abra o Terminal**.

2. **Edite o Arquivo de Configuração do Shell**:

   - Para `bash`, edite o arquivo `~/.bash_profile` ou `~/.bashrc`.
   - Para `zsh`, edite o arquivo `~/.zshrc`.

3. **Adicione o Caminho do SDK**:

   - Adicione as seguintes linhas ao final do arquivo:
     ```sh
     export ANDROID_HOME=/Users/<SeuUsuário>/Library/Android/sdk
     export PATH=$PATH:$ANDROID_HOME/platform-tools
     ```

4. **Salve e Feche** o arquivo.

5. **Atualize as Variáveis de Ambiente**:
   - No terminal, execute:
     ```sh
     source ~/.bash_profile
     ```
     ou, para `zsh`:
     ```sh
     source ~/.zshrc
     ```

### 4. Verificando a Configuração

Para verificar se as variáveis de ambiente estão configuradas corretamente, abra um novo terminal e digite:

```bash
adb version
```
