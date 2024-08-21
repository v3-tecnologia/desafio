# Desafio Desenvolvedor Backend Pleno - V3

Desafio consiste em coletar métricas de dispositivos embarcados.

# Requisitos

Algumas ferramentas são necessárias para utilizar a aplicação:

- Docker

# Execução

O ambiente de execução é inserido em um Docker Compose, onde possui tanto os serviços de banco de dados Postgres quanto a própria API Golang. Para executar, clone o repositório e execute os comandos:

```
$ docker compose up --build -d
```

# Tecnologias e Bibliotecas

- Golang 1.22 - Linguagem de programação
- PostgreSQL - Banco de dados
- SQLX - Executar operações no banco de dados
- Golang Migrate - Executar migrações no banco de dados
- Testify - Realizar testes mais claros
- AWS Rekognition

# API

- Url base: localhost:3000

## Rotas

| Método | URL                  | Privada | Funcionalidade               |
| :----- | -------------------- | :-----: | ---------------------------- |
| POST   | /telemetry/photo     |    -    | Cria dados de uma foto.      |
| POST   | /telemetry/gps       |    -    | Cria dados de um GPS.        |
| POST   | /telemetry/gyroscope |    -    | Cria dados de um giroscópio. |

## Dados para execução

### `[POST]` /telemetry/gps

Caso de Sucesso: Rota para criar dados de um gps.

Body da requisição:

```
{
    "latitude": 28.99,
    "longitude": -491.99
}
```

### `[POST]` /telemetry/gyroscope

Caso de Sucesso: Rota para criar dados de um giroscópio.

Body da requisição:

```
{
    "x": -192.4,
    "y": 99.0,
    "z": 12.944
}
```

### `[POST]` /telemetry/photo

Caso de Sucesso: Rota para criar dados de uma foto.

- Em um Multipart Form:
  - Criar chave com nome: "photo" e anexar uma imagem com formato válido.
