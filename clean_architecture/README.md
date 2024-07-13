## Desafio de Clean Architecture
Desafio do do módulo de Clean Architecture.

### Objetivo
Necessário criar o usecase de listagem de `orders` que disponibilize a consulta de ordens através das interfaces Http, GraphQL e GRPC.

### Requisitos
As interfaces de listagem precisam ser:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

### Executando a aplicação
Acesse o diretório `clean_architecture/` e execute o comando `docker-compose up` para instanciar os componentes de infraestrutura.

Acesse o diretório `cmd/` e execute o comando `go run main.go wire_gen.go` e os logs de boot da aplicação serão exibidos.

### Criando uma ordem

**HTTP**
```
curl --location --request POST 'http://localhost:8000/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"a4c9fh",
    "price": 954,
    "tax": 934
}'
```

**GraphQL**

Acesse a URL `http://localhost:8080/` e execute a mutation:

```
mutation createOrder {
  createOrder(input:{
      id:"zzsa", Price:777, Tax:3
  }){
    id Price Tax FinalPrice
  }
}
```

**GRPC**

Execute o comando do Evans CLI `evans repl -r` e após conectar no localhost `127.0.0.1:50051`

```
- Selecione o pacote `package pb`
- Selecione o serviço `service OrderService`
- Selecione a chamada `call CreateOrder`
- Informe os dados da chamada e receba o response
```

### Listando todas as ordens

**HTTP**

```
curl --location --request GET 'http://localhost:8000/orders' \
--data-raw ''
```

**GraphQL**

Acesse a URL `http://localhost:8080/` e execute a query:

```
query {
  orders{id, Tax, Price, FinalPrice}
}
```

**GRPC**

Execute o comando do Evans CLI `evans repl -r` e após conectar no localhost `127.0.0.1:50051`

```
- Selecione o pacote `package pb`
- Selecione o serviço `service OrderService`
- Selecione a chamada `call FetchAllOrders`
- Informe os dados da chamada e receba o response
```