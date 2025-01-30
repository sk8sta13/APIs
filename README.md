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

### REST

Para rodar o endpoint REST é só executar o comando:

```
go run cmd/rest/main.go
```

E executar a request configurada no arquivo "test/order.http". Veja um print do exemplo abaixo:

![image](https://github.com/user-attachments/assets/79b40759-f91f-4dab-97a4-ae9393994b91)

### gRPC

Para rodar o server gRPC execute o comando:

```
go run cmd/grpc/main.go
```

Para executar os testes estou utilizando o Evans:

```
docker run --rm -it --network host -v "$(pwd):/mount:ro" ghcr.io/ktr0731/evans:latest --path ./proto --proto settings.proto --host localhost --port 50051 repl
```

Ao carregar a tela do Evans, digite o comando:

```
call ListOrders
```

Veja o exemplo no print abaixo:

![image](https://github.com/user-attachments/assets/0afb6362-c57a-40fc-a2f4-c5e559fb344d)
