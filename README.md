# MVP-Telemetria-V3

Este repositório contém a implementação do desafio Backend da V3, em Golang. Ele contempla um MVP para uma venda a um cliente fictício, que recebe dados de telemetria de um celular Android.


## Índice

- [Desafio](#Desafio)
- [Arquitetura](#arquitetura)
- [Bancos de dados](#bancos de dados)
- [Instalação e execução](#Instalação e execução)
- [Endpoints REST](#endpoints-rest)
- [Testes](#testes)

## Desafio
As informações coletadas são:

1. **Dados de Giroscópio** 
2. **Dados de GPS** 
3. **Uma foto**

Este desafio foi implementado até o nível 4, na parte de containerização.

## Arquitetura

Para este desafio, foi pensada a utilização da implementação de chamadas por casos de uso, pensando nos conceitos de SOLID, e também facilitando no entendimenton de cada etapa do projeto. 

## Bancos de dados

Foi escolhido o trabalho com o banco de dados relacional, tendo em vista que todas as informações enviadas, tem um ponton em comum:
O dispositivo, que será identificado pelo MacAddress. O banco relacional facilida para manter os dados enviados pelo mesmo dispositivo, ligados de forma concisa.

o modelo utilizado foi o que está abaixo representado:

![database-diagram.png](resources%2Fdatabase-diagram.png)


## Instalação e execução

Clone o repositório para sua máquina local:

```bash
git clone https://github.com/icaromarques/desafio-backend.git
```

Execute o projeto utilizando o Docker compose

```bash
cd desafio-backend
docker compose up
```

## Endpoints REST

### Dados enviados por GPS

`POST: /telemetry/gps`
#### Request:

```json
{
 "timestamp" : "2024-12-26T18:00:05Z",
 "macAddress" : "99:98:ca:89:be:ce",
 "latitude": 4.245422,
 "longitude": 3.233254
}
```


#### Response:

```json
{
  "macAddress": "99:98:ca:89:be:ce",
  "timestamp": "2024-12-26T18:00:05Z",
  "coordinates": "{\"type\":\"Point\",\"coordinates\":[4.245422,3.233254]}"
}
```


### Dados enviados por Giroscópio

`POST: /telemetry/gyroscope`
#### Request:

```json
{
  "timestamp" : "2024-08-26T18:00:05Z",
  "macAddress" : "99:98:ca:89:be:fe",
  "xAxis": 0.245422,
  "yAxis": 1.233254,
  "zAxis": 0.213234
}
```


#### Response:

```json
{
    "macAddress": "99:98:ca:89:be:fe",
    "timestamp": "2024-08-26T18:00:05Z",
    "xAxis": 0.245422,
    "yAxis": 1.233254,
    "zAxis": 0.213234
}
```


### Dados enviados de fotos

`POST: /telemetry/photo`
#### Request:

```json
request (type= text) :{
  "timestamp" : "2024-12-26T18:00:05Z",
 "macAddress" : "99:98:ca:89:be:fe"
}

immage (type= file): selected image 
```


#### Response:

```json
{
  "macAddress": "99:98:ca:89:be:fe",
  "timestamp": "2024-12-26T18:00:05Z",
  "latitude": {
    "Name": "uploaded-image-2547156884.jpg",
    "FullPath": "C:\\Users\\Zeus\\Documents\\workspace\\desafio-backend\\resources\\temp-files\\uploaded-image-2547156884.jpg",
    "MimeType": "",
    "Bytes": null
  }
}
```

obs.: no diretório 
```resources``` encontra-se a collection Postman que pode ser usada para testar todos os endpoints.  
## Testes

Para rodar os testes unitários do projeto, utilize o comando:

```bash
go test ./...
```

