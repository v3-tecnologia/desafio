# Desafio Técnico V3

Este repositório consiste em uma aplicação backend desenvolvida em GO para recebimento e salvamento de dados enviados por um equipamento através de um aplicativo. O objetivo do desafio é criar uma API que receba dados de giroscópio, GPS e fotos, validá-los e salvar essas informações em banco de dados, além de realizar testes unitários para garantir a funcionalidade da aplicação.

## API de recebimeto de dados

Este projeto consiste em uma API para receber e validar dados de equipamentos de telemetria, que possui endpoints para processamento e armazenamento de dados.

## Requisitos do projeto
- Go 1.23
- Docker e docker-compose
- PostgreSQL

## Estrutura dos diretórios do projeto
- handlers: contém as funções e testes referentes às requisições.

- migration: funções de criação das tabelas do banco de dados.

- models: contém as structs básicas utilizadas no projeto

- repository: contém as funções relacionadas ao banco de dados.

- routes: configuração das rotas utilizadas pela API.

- service: funções para manipulação de dados e regras do serviço.

## Endpoints
### 1. Dados de giroscópio (POST /telemetry/gyroscope)
- Recebe e valida os dados do giroscópio
- O corpo da requisição deve conter o json:

  ```json
   {
        "macAddress": "00:A6:C7:87:F7:26",
        "x": 123.1, 
        "y": 213.2,
        "z": 32.3, 
        "timestamp": 1724855500 
    }

### 2. Dados de GPS (POST /telemetry/gyroscope)
- Recebe e valida os dados do giroscópio
- O corpo da requisição deve conter o json:

  ```json
   {
        "macAddress": "00:A6:C7:87:F7:26",
        "latitude": "-5.088889", 
        "longitude": "-42.801944",
        "timestamp": 1724855500 
    }

### 3. Dados de foto (POST /telemetry/gyroscope)
- Recebe e valida os dados do giroscópio
- O corpo da requisição deve conter o json:

  ```json
   {
        "macAddress": "00:A6:C7:87:F7:26",
        "photo": "/9j/4AAQSkZJRgABAQEBLAEsAAD ...", 
        "timestamp": 1724855500 
    }

## Armazenamento de dados

Os dados são armazenados em um banco de dados PostgreSQL. O banco de dados foi configurado usando o docker-compose, que executa uma instância pré-configurada com as tabelas usadas no projeto. Foi implementada uma migration para a criação das tabelas gyroscope_data, gps_data e photo.

## Testes

Os testes unitários foram realizados em funções essenciais para a aplicação, como as funções do handler e service.

### Ferramentas de teste
- Testing
- [Testify](github.com/stretchr/testify)
- [Mockery](github.com/vektra/mockery)

A utilização do pacote mockery foi necessário para isolar o comportamento das interfaces da comunicação de ferramentas externas, por meio da geração de mocks.

Para executar os testes, use o comando:
``` 
go test -v -cover ./...
```
No makefile criado para a execução da aplicação, existe comandos para a criação dos mocks e rodar os testes:
``` 
make build-mocks
```
``` 
make run-tests
```

## Execução da aplicação

O repositório possui um makefile em que há os comandos para testes e execução da aplicação. Isso foi feito para facilitar a utilização dos comandos. Para a aplicação, comando é:
``` 
make run-app
```
Este comando executa os contêineres configurados no arquivo docker-compose.yml

## Considerações finais

O projeto foi desenvolvido a partir das diretrizes propostas no desafio. Foram realizadas todas as atividades até o nível 4. Não foi possível realizar o nível 5, pois para atender o prazo estabelecido, não houve tempo hábil para completar o que foi proposto nesse nível.