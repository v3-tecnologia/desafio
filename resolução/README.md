# Desafio Backend em Go

Este repositório contém uma aplicação backend desenvolvida em Go que recebe e processa dados enviados por um aplicativo. O objetivo é criar uma API que receba dados de giroscópio, GPS e fotos, valide essas informações, armazene-as em um banco de dados e forneça testes unitários para garantir a funcionalidade correta.

## Sumário
1. [Visão Geral]()
2. [Requisitos]()
3. [Endpoints da API]()
4. [Banco de Dados]()
5. [Testes]()
6. [Como Executar]()
7. [Como Testar]()


## Visão Geral

Esta aplicação é um serviço backend desenvolvido em Go que oferece uma API para receber e validar dados. O sistema é projetado para ser expandido com testes unitários e containerização com Docker.

## Requisitos
- Go 1.23
- Docker
- Docker Compose (opcional, para configuração de ambiente)

## Endpoints da API

A API possui os seguintes endpoints:

1. **POST /telemetry/gyroscope**

Recebe dados do giroscópio. O corpo da requisição deve conter:

     ```json
    {
     "x": "float",
     "y": "float",
     "z": "float"
    }

3. **POST /telemetry/gps**

Recebe dados de GPS. O corpo da requisição deve conter:

    {
        "latitude": "float",
        "longitude": "float",
        "altitude": "float"
    }

4. **POST /telemetry/photo**

Recebe dados da foto. O corpo da requisição deve conter:

    {
        "photo_id": "string",
        "photo_data": "base64_encoded_string"
    }

## Banco de Dados

Essa aplicação foi feita para armazenar os dados em um banco de dados Postgresql, para esse desafio ao [executar o projeto]() uma instância pré-configurada será executada contendo as tabelas gyroscope, gps e photo para armazenamento dos dados

## Testes

Essa aplicação contém testes unitarios para todas as funções que tem a necessidade. Nesse caso as funções de handler e de serviço. Ao [executar o projeto]() existe uma validação dos casos de teste durante o build, caso os testes falhem o build não é concluido.
Para execução dos teste via terminal use algum dos comandos: 

    github.com/vektra/mockery                                                           
    go test -v -cover ./...

### Ferramentas de Teste

    testing
    github.com/stretchr/testify
    github.com/vektra/mockery

## Docker

Esse repositorio conta com um Dockerfile e um docker-compose.yml para configuração e utilização dos containers contendo a aplicação e o banco de dados.

## Como Executar

Clone o Repositório

bash

git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>

Construa e Inicie os Containers

bash

    docker-compose up --build

    Isso construirá a imagem Docker da aplicação e iniciará o serviço junto com o banco de dados.

Como Testar

    Executar Testes Unitários

    Dentro do container da aplicação, execute:

    bash

go test ./...

Testar a API

Você pode usar ferramentas como curl ou Postman para testar os endpoints da API:

bash

    curl -X POST http://localhost:8080/telemetry/gyroscope -d '{"x": 1.23, "y": 4.56, "z": 7.89}' -H "Content-Type: application/json"

Contribuição

Sinta-se à vontade para contribuir com o projeto! Para isso:

    Faça um fork do repositório.
    Crie uma branch para suas alterações.
    Envie um pull request com uma descrição clara das mudanças.