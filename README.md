# Desafio Técnico V3

## Setup
A aplicação está dockerizada, desde que sua porta 3000 esteja livre você pode facilmente rodar `docker-compose up --build` e você terá a aplicação rodando. Se não, basta mudar a porta sendo exposta na configuração da aplicação.

## Testando
Para testar basta navegar até a pasta de testes que você quer executar pelo terminal e rodar `go test`

## Usando a aplicação
Navegue até http://localhost:3000/swagger/index.html e você terá a documentação do swagger (assumindo que você não modificou a porta em que a aplicação está sendo servida), que vai te mostrar os 3 endpoints POST disponíveis.

## O problema proposto
Não tenho conhecimento de como é a codebase da V3, mas imagino que deva ser infinitamente mais complexa do que o que eu apresento aqui.

Esta arquitetura e os patterns que eu optei por usar aqui são somente para resolver este problema proposto, que tem um escopo limitado.

Esta arquitetura serve muito bem para uma API comum que ordinariamente se preocuparia somente com CRUDs com várias regras de negócio, podendo servir até alguns milhões de requests por dia em prod. Outra vantagem é que ela pode ser facilmente convertida para serverless com Lambda ou qualquer outra solução cloud.

### Os problemas 
As regras de negócio atualmente vivem nos modelos.Para o escopo deste teste, isso não importa, mas facilmente poderia fugir de controle se tivessemos algo mais complexo em mãos.


## Até onde eu fiz o desafio
Eu fiz o desafio backend, até o nível 4.

### Features
<ul>
  <li>Testes unitários pra cada modelo</li>
  <li>Teste de integração utilizando a API pública</li>
  <li>Middleware pra Throttling e rate limiting</li>
  <li>Containerização</li>
  <li>Arquitetura simples e comummente usada</li>
  <li>Middleware de logging</li>
  <li>Erros customizados</li>
</ul>

## Arrependimentos e incertezas
Eu imagino que para um teste armazenamento na memória está ok, mas talvez eu devesse ter incluido uma DB de verdade, o que seria trivial de incluir no container e na arquitetura atual.

Não incluí nem a configuração do Rekognition, pois não tenho certeza de como ele seria implementado, então optei por não fazer, escolhendo fazer menos porém de maneira certa. Se eu fosse responsável por implementar isso em prod eu provavelmente poderia fazer sem muitos problemas, já que a documentação da AWS é bem extensiva, mas este seria outro cenário.