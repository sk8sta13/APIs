# APIs

Esse projeto é relativo a um dos desafios do curso "Pós Go Expert - 2024", consiste em um banco de dados com uma tabela chamada "orders" onde temos uma massa de dados apenas para testes, e como APIs temos três formas uma é utilizando o padrão REST outro vamos utilizar o GraphQL e por fim vamos ter um também no padrão gRPC.

Todas essas APIs são feitas utilizando Go.

### Instalação

Clone esse repositório e em seguida entre na pasta do projeto e la dentro execute o comando:

```
docker compose up -d
```

Isso fara com que um container com o banco dedados MySQL seja criado e também seja feito o dump de uma pequena quantidade de dados apenas para testes.

Para parar o container execute o stop:

```
docker compose stop
```