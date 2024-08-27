# Desafio tecnico V3

Este projeto consiste em uma api para cadastro de informações sobre itens de um determinado device (gps, gyroscope e photo).

## Informação

- **Versão**: 1.0
- **Host**: localhost:5000
- **Etapas**: Foi implementado até a 4 etapa.

## Pré-requisitos

Para rodar o projeto é necessário:

- [Go](https://golang.org/dl/): Linguagem utilizada no projeto.
- [Docker](https://www.docker.com/get-started): Caso queira rodar em container.
- [PostgreSQL](https://www.postgresql.org/): Banco de dados utilizado.

## Instalação

1. **Clone o repositório:**
2. **Popular banco de dados:**
    ```
    CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    mac_address character varying UNIQUE NOT NULL
    );

    CREATE TABLE gyroscopes (
    id SERIAL PRIMARY KEY,
    device_id INTEGER NOT NULL REFERENCES devices(id),
    x NUMERIC NOT NULL,
    y NUMERIC PRECISION NOT NULL,
    z NUMERIC PRECISION NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );

    CREATE TABLE gps (
    id SERIAL PRIMARY KEY,
    device_id INTEGER NOT NULL REFERENCES devices(id),
    latitude NUMERIC NOT NULL,
    longitude NUMERIC NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );

    CREATE TABLE photos (
    id SERIAL PRIMARY KEY,
    device_id INTEGER NOT NULL REFERENCES devices(id),
    url CHARACTER VARYING NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );

    insert into public.devices(mac_address) VALUES ('4A-66-45-37-8D-05');
    ```

4. **Navegue até a pasta do projeto:**

   ```
   cd desafio
   ```

5. **Build a aplicação usando Docker Compose:**

   ```
   docker compose up
   ```

## Rode a aplicação

Depois da instalação, você pode rodar o projeto em Go com o seguinte comando (se você preferir pode rodar direatemente na IDE):

```
cd cmd
cd api
go run main.go
```

A aplicação ficará acessivel em `http://localhost:5000`.

## Teste a aplicação

- **Testes automatizados:**
  ```
  cd internal
  cd service
  go test
  ```

  ## Curls das requests

- **Criar gyroscope:**

  ```
   curl -X POST 'localhost:5000/telemetry/gyroscope' -H 'Content-Type: application/json' -d '{
    "x": 10,
    "y": 10,
    "z": 12,
    "mac_address": "4A-66-45-37-8D-05",
    "collection_date": "2024-08-26T16:35:54.123456Z"
  }'
  ```

- **Criar GPS:**

  ```
  curl -X POST 'localhost:5000/telemetry/gps' -H 'Content-Type: application/json' -d '{
    "latitude": 12,
    "longitude": 14,
    "mac_address": "4A-66-45-37-8D-05",
    "collection_date": "2024-08-26T16:35:54.123456Z"
  }'
   ```

- **Criar photo:**

  ```
  curl -X POST 'localhost:5000/telemetry/photo' -H 'Content-Type: application/json' -d '{
    "url": "www.teste.com.br",
    "mac_address": "4A-66-45-37-8D-05",
    "collection_date": "2024-08-26T16:35:54.123456Z"
  }'
   ```
