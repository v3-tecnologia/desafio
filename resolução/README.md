# Desafio Backend em Go

Este repositório contém uma aplicação backend desenvolvida em Go que recebe e processa dados enviados por um aplicativo. O objetivo é criar uma API que receba dados de giroscópio, GPS e fotos, valide essas informações, armazene-as em um banco de dados e forneça testes unitários para garantir a funcionalidade correta.

## Sumário
1. [Visão Geral](#visão-geral)
2. [Requisitos](#requisitos)
3. [Endpoints da API](#endpoints-da-api)
4. [Banco de Dados](#banco-de-dados)
5. [Testes](#testes)
6. [Como Executar](#como-executar)
8. [Diretorios](#diretorios)
7. [Considerações Finais](#considerações-finais)


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
        "macAddr" : "00-B0-D0-63-C2-26", //string mac address
        "x" : 111.2, //float coordenada x
        "y" : 222.3, //float coordenada y
        "z" : 333.4, //float coordenada z
        "timeStamp" : 1724603773 //int timestamp em formato UNIX
    }

2. **POST /telemetry/gps**

Recebe dados de GPS. O corpo da requisição deve conter:

    {
        "macAddr" : "00-B0-D0-63-C2-26", //string mac address
        "latitude" : "-18.909762",  //string latitude
        "longitude" : "-48.232750", //string longitude
        "timeStamp" : 1724603773 //int timestamp em formato UNIX
    }

3. **POST /telemetry/photo**

Recebe dados da foto. O corpo da requisição deve conter:

    {
        "macAddr" : "00-B0-D0-63-C2-26", //string mac address
        "timeStamp" : 1724603773, //int timestamp em formato UNIX
        "image" : "/9j/4AAQSkZJRgABAQAAAQABAAD/..." //string imagem em base64
    }

## Banco de Dados

Essa aplicação foi feita para armazenar os dados em um banco de dados Postgresql, para esse desafio ao [executar o projeto](#como-executar) uma instância pré-configurada será executada contendo as tabelas gyroscope, gps e photo para armazenamento dos dados

## Testes

Essa aplicação contém testes unitarios para todas as funções que tem a necessidade. Nesse caso as funções de handler e de serviço. Ao [executar o projeto](#como-executar) existe uma validação dos casos de teste durante o build, caso os testes falhem o build não é concluido.
Para execução dos teste via terminal use o comando: 
                                                
    go test -v -cover ./...

Alternativamente esse diretorio possui um makefile com os comandos **build-mocks** que gera os mocks e depencias das ferramentas de teste e o comando **tests** que executa o comando citado a acima.

### Ferramentas de Teste

    testing
    github.com/stretchr/testify
    github.com/vektra/mockery

## obs
Vale destacar que foi utilizado o pacote mockery para gerar mocks de interface permitindo utilizar isolar o comportamento das funções que precisariam se comunicar com ferramentas externas.

## Docker

Esse repositorio conta com um Dockerfile e um docker-compose.yml para configuração e utilização dos containers contendo a aplicação e o banco de dados.

## Como Executar

Esse Repositório possui um makefile para facilitar a excução.
O comando **make run** builda e inicia os containers.

## Diretorios

- handlers
    - Contém as funções e testes para lidar com as requisições
- models
    - Contém os modelos de structs essenciais que serão usados por todo código
- repository
    - Contém funções para conexões externas e suas interações
- routes
    - Contém as configurações de rotas so servidor utilizando gorillaMux
- service
    -   Contém as funções de regras de serviços e manipulações de dados

## Considerações Finais
Todo esse projeto foi construido com base nas diretrizes do desafio e foram realizadas todas as atividades até o nível 4, referente ao nível 5 um modelo foi treinado e pode ser disponibilizado para avaliação caso necessário, porém por falta de familiaridade não foi possivél realizar as integrações da maneira portanto ficaram de fora.