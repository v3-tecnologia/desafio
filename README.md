# Telemetry Microservice

Este projeto é um microserviço de telemetria desenvolvido em Go, seguindo os princípios de Clean Architecture,
utilizando PostgreSQL como banco de dados.

## Descrição

O Telemetry Microservice é responsável por coletar e gerenciar dados de telemetria de dispositivos em um sistema
web. Ele fornece uma API RESTful para receber e armazenar dados de giroscópio, GPS e fotos.

## Tecnologias Utilizadas

- Go (Golang)
- PostgreSQL
- Docker
- Swagger (para documentação da API)

## Arquitetura

Este projeto segue os princípios de Clean Architecture e Domain-Driven Design:

- **Domain Layer**: Contém as entidades de negócio e regras de domínio.
- **Use Case Layer**: Implementa a lógica de aplicação e casos de uso.
- **Infrastructure Layer**: Fornece implementações concretas para interfaces definidas em camadas superiores.

## Funcionalidades

O microserviço oferece os seguintes endpoints:

1. `POST /telemetry/gyroscope`: Recebe dados do giroscópio (x, y, z).
2. `POST /telemetry/gps`: Recebe dados de GPS (latitude, longitude).
3. `POST /telemetry/photo`: Recebe uma foto.

Todos os endpoints requerem um campo com a identificação única do dispositivo (endereço MAC).

> **NOTA:** O TIMESTAMP não é recebido, mas é gerado no momento da request pela aplicação.

## Validação de Dados

O microserviço garante que todos os dados recebidos estejam preenchidos corretamente. Caso algum dado esteja faltando,
uma mensagem de erro é retornada com o Status 400.

## Armazenamento de Dados

Todos os dados são armazenados no banco de dados PostgreSQL junto com o timestamp do momento em que foram coletados.

> **NOTA:** Para o endpoint de fotos (/telemetry/photo), os arquivos/imagens são salvos no diretório 'uploads' dentro da
> aplicação. No banco de dados, apenas o caminho do arquivo é armazenado, seguindo uma prática comum de gestão de
> arquivos. Em um cenário de produção real, essas imagens deveriam ser armazenadas em um repositório ou bucket na nuvem
> para melhor escalabilidade e gerenciamento, mantendo a prática de armazenar apenas o caminho ou identificador do arquivo
> no banco de dados.

## Como Executar

Para executar este projeto, siga os passos abaixo:

1. Certifique-se de ter o Docker instalados em sua máquina.

2. Clone o repositório:

   `git clone https://github.com/HaroldoFV/desafio`

   `cd desafio`

3. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```
    DB_DRIVER=postgres
    DB_HOST=postgres
    DB_PORT=5432
    DB_USER=seu_usuario
    DB_PASSWORD=sua_senha
    DB_NAME=nome_do_banco
    WEB_SERVER_PORT=8080
    TEST_DB_HOST=localhost
    TEST_DB_PORT=5433
    TEST_DB_USER=seu_usuario
    TEST_DB_PASSWORD=sua_senha
    TEST_DB_NAME=nome_do_banco_test

> **NOTA:** Para facilitar os testes adicionei o arquivo .env ao repositório(não se deve versionar esse tipo
> de arquivo)

4. Inicie os serviços usando Docker Compose:

   `docker-compose up -d`

   Isso irá iniciar o banco de dados PostgreSQL, executar as migrações e iniciar a aplicação.

5. A aplicação estará disponível em `http://localhost:8080/docs/index.html`.

## Testes

Para executar os testes, siga estas etapas:

1. Certifique-se de que o container de teste do PostgreSQL(telemetry_db_test) está em execução:

   `docker-compose ps`

2. Execute os testes usando o seguinte comando:

   `go test ./... -v`

Este comando executará todos os testes no projeto, incluindo testes de unidade e integração.

> **NOTA:** Os testes de integração usarão o banco de dados de teste (postgres_test) que está configurado para rodar na
> porta 5433.

## Testando as Requisições

Para facilitar o teste das requisições da API, fornecemos arquivos .http na pasta `api`. Esses arquivos podem ser usados
para testar as requisições diretamente de IDEs compatíveis (como VSCode com a extensão REST Client) ou podem ser
convertidos para cURL ou outras ferramentas de sua preferência.

Os arquivos .http incluem exemplos de requisições para cada endpoint da API, permitindo que você teste rapidamente a
funcionalidade do microserviço.