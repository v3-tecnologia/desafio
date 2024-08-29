# Desafio Técnico V3

# Backend

Objetivo: Criar um back-end em Go para receber e armazenar os dados enviados pelo aplicativo mobile.
O desafio é dividido em níveis.

## Nível 1

### Tarefa

Deve-se criar uma API que receba requisições de acordo com os endpoints:

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

Deve-se garantir que os dados recebidos estão preenchidos corretamente.

Caso algum dado esteja faltando, então retorne uma mensagem de erro e um Status 400.

### Solução

Foi feito uma aplicação em Go com os endpoints descritos onde os dados são recebidos em formato JSON.
Em todos os endpoint, todos os campos são obrigatórios. O deviceID é uma string qualquer utilizada para identificar um dispositivo, endereço MAC ou outro.
O timestamp esta no formato unix, como um numero de 64 bits que é o numero de segundos desde 01/01/1970.

Para o endpoint `POST /telemetry/gyroscope` o corpo do pacote é:

```
{
  "deviceID": string,
  "timestamp": uint,
  "x": float,
  "y": float,
  "z": float
}
```

Para o endpoint `POST /telemetry/gps` o corpos do pacote é:

```
{
  "deviceID": string,
  "timestamp": uint,
  "logitude": float,
  "latitude": float
}
```

Para o endpoint `POST /telemetry/photo` o corpo do pacote é:

```
{
  "deviceID": string,
  "timestamp": uint,
  "image": string,
  "format":string 
}
```

Neste endpoint a imagem é recebida como o blob da imagem codificada com base64.
O formato é em que a imagem foi codificada (jpg, png, bmp, etc).

Instruções de como rodar a aplicação estão após a descrição do banco de dados.

## Nível 2

### Tarefa

Salve cada uma das informações em um banco de dados a sua escolha.

Salve estes dados de forma identificável e consistente;

### Solução

O banco da aplicação é um mysql organizado em 4 tabelas:
  - devices;
  - gyroscope;
  - gps;
  - photos.

A tabela devices é definida como no sql abaixo e é utilizada para validar os IDs recebidos.
Esta tabela também facilita possiveis expansões futuras do sistema.

```
devices (
    deviceID varchar(255) not null primary key
    );
```

A tabela gyroscope guarda as informações do giroscópio e é definida pelo sql abaixo.
A chave primária é o identificador do dispositivo e o timestamp da captura. Portanto, dados do dispositivos deverão ser capturados com intervalos maiores que 1 segundo.
O identificador do dispositivo faz referência a tabela devices e só são aceitos dados de dispositivos cadastrados.
As restições de chave primária e estrangeira são replicadas para as tabelas gps e photos.

```
gyroscope (
    deviceID varchar(255) not null,
    x float(24) not null,
    y float(24) not null,
    z float(24) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint gyroscope_device foreign key (deviceID) references devices(deviceID)
    );
```

A tabela gps guarda as informações do giroscópio e é definida pelo sql abaixo.

```
gps (
    deviceID varchar(255) not null,
    latitude float(24) not null,
    longitude float(24) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint gps_device foreign key (deviceID) references devices(deviceID)
    );
```

A tabela photos guarda as informações relacionadas as fotos e é definida pelo sql abaixo.
A tabela não guarda a foto propriamente dita, mas o caminho para onde a foto está salva no servidor.
Para evitar inconsistencias com blobs no banco de dados, as fotos são guardas em um diretório no servidor.
O nome do arquivo utilizado para cada foto é o mesmo da chave primária da foto no banco de dados seguido do formato salvo. Ou seja, <deviceID>-<timestamp>.<formato>

```
photos (
    deviceID varchar(255) not null,
    photo varchar(255) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint photo_device foreign key (deviceID) references devices(deviceID)
    );
```

### Executando a aplicação

Para executar a aplicação são necessárias 3 preparações do ambiente:
  - Configurar o banco de dados;
  - Criar um arquivo de configuração do acesso ao banco;
  - Ter um diretório para armazenar as fotos.

#### Configuração do banco

Os passo são:
 - Com o mysql-server instalado, acesse o banco com o comando "mysql -u root -p";
 - Depois crie um banco chamado "v3";
 - Por fim execute o script "create_devices.sql" do diretório "database".

O script cria um usuário para a aplicação com o acesso ao banco "v3", cria todas as tabelas que ainda não foram criadas e cadastra alguns dispositivos para demostração.
Os dispositivos são cadastrados com identificadores Device1, Device2 e Device3.

#### Arquivo de configuração do acesso

O arquivo consiste em uma sequência de campos em texto pleno com cada campo em uma linha separada. O arquivo precisa conter as seguintes informações:
  - Usuário;
  - Senha;
  - Protocolo de acesso (tcp ou udp);
  - Endereço do banco e porta;
  - Nome do banco que será acessado.

Os campos do arquivo precisam estar nesta ordem especificamente.

Considerando um banco rodando no mesmo servidor da aplicação com configurações padrão e gerado com o script "create_devices.sql".
Um exemplo de arquivo de configuração está disponível na pasta "app" com o nome "database.cfg".

#### Rodando a aplicação

A aplicação começa a rodar após a chamada do comando "go run . <nome do arquivo de configuração> <caminho para o diretório de fotos>".

## Nível 3

### tarefa

Crie testes unitários para cada arquivo da aplicação. Para cada nova implementação a seguir, também deve-se criar os testes.

### Solução

Para cada arquivo <nome>.go da aplicação foi feito um arquivo <nome>_test.go contendo os testes para as funções do arquivo.
Alguns testes testam a interação entre a aplicação e o banco de dados. Portanto, para a execução dos testes é necessário um banco local com um usuário 'tester'@'localhost' com senha 'test_passwd' e acesso pleno a um banco 'db_test'.

Com o usuário do banco criado, os testes podem ser rodados com o comando "go test".

## Nível 4

### tarefa

Crie um _container_ em _Docker_ que contenha a sua aplicação e o banco de dados utilizado nos testes.

### solução

Foram criados dois dockerfiles um para a aplicação e um para o banco de dados.
O dockerfile do banco está no diretório "database". Para gerar a imagem, basta rodar o comando "docker build -t database ." dentro do diretório.
O dockerfile da aplicação está no diretório "app". A geração da imagem de aplicação utiliza o arquivo "database.cfg". Ajuste o enderço no arquivo para endereço onde o banco esta hospedado, se necessário. Depois execute o comando "docker build -t app ." na pasta "app".

Para rodar a aplicação execute os comandos:

```
mkdir db_data
mkdir photo_data
docker run --name database --network="host" -v ./db_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root_passwd -d database_desafio
docker run --name application -v ./photo_data:/usr/app/images --network="host" app_desafio
```

## Nível 5

### tarefa
A cada foto recebida, deve-se utilizar o AWS Rekognition para comparar se a foto enviada é reconhecida com base nas fotos anteriores enviadas.

Se a foto enviada for reconhecida, retorne como resposta do `POST` um atributo que indique isso.

Utilize as fotos iniciais para realizar o treinamento da IA.

### solução

Não consegui realizar está tarefa a tempo.