# API REST de Clientes em Go

API REST desenvolvida em **Go (Golang)** utilizando o pacote **Gorilla Mux**.

A aplicação permite cadastrar clientes, listar clientes e buscar um cliente específico pelo seu ID.

Os dados são armazenados temporariamente em memória durante a execução da aplicação.

## Tecnologias utilizadas

- Go (Golang)
- Gorilla Mux
- HTTP
- JSON
- Postman

## Estrutura do projeto

```text
api-go-clients/
│
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Como executar o projeto

### 1. Pré-requisitos

É necessário ter o Go instalado no computador.

Para verificar se o Go está instalado, execute:

```bash
go version
```

### 2. Acessar a pasta do projeto

Abra o terminal na pasta principal do projeto:

```bash
cd api-go-clients
```

> **Importante:** execute os comandos a partir da pasta onde está localizado o arquivo `go.mod`.

### 3. Instalar as dependências

O projeto utiliza o pacote **Gorilla Mux**.

Caso ainda não esteja instalado, execute:

```bash
go get -u github.com/gorilla/mux
```

Depois:

```bash
go mod tidy
```

### 4. Executar a aplicação

Na pasta principal do projeto, execute:

```bash
go run .
```

O servidor será iniciado na porta `8080`.

A API estará disponível em:

```text
http://localhost:8080
```

Mantenha o terminal aberto enquanto estiver utilizando a API.

Para encerrar o servidor, pressione:

```text
Ctrl + C
```

## Endpoints

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/clientes` | Lista todos os clientes |
| POST | `/clientes` | Cadastra um novo cliente |
| GET | `/clientes/{id}` | Busca um cliente pelo ID |

## Testando a API

### 1. Listar clientes

Utilize:

```text
GET http://localhost:8080/clientes
```

Caso nenhum cliente tenha sido cadastrado, a resposta será:

```json
[]
```

### 2. Cadastrar um cliente

No Postman, crie uma nova requisição utilizando o método:

```text
POST
```

URL:

```text
http://localhost:8080/clientes
```

Depois, acesse:

**Body → raw → JSON**

Insira:

```json
{
    "nome": "João Silva",
    "email": "joao@email.com"
}
```

Clique em **Send**.

A resposta será:

```json
{
    "id": 1,
    "nome": "João Silva",
    "email": "joao@email.com"
}
```

O ID é gerado automaticamente pela aplicação.

### 3. Cadastrar outro cliente

Utilize novamente:

```text
POST http://localhost:8080/clientes
```

Com o seguinte JSON:

```json
{
    "nome": "Maria Souza",
    "email": "maria@email.com"
}
```

A resposta será:

```json
{
    "id": 2,
    "nome": "Maria Souza",
    "email": "maria@email.com"
}
```

### 4. Listar todos os clientes

Faça uma requisição:

```text
GET http://localhost:8080/clientes
```

A resposta será semelhante a:

```json
[
    {
        "id": 1,
        "nome": "João Silva",
        "email": "joao@email.com"
    },
    {
        "id": 2,
        "nome": "Maria Souza",
        "email": "maria@email.com"
    }
]
```

### 5. Buscar um cliente pelo ID

Para buscar um cliente específico, utilize:

```text
GET http://localhost:8080/clientes/1
```

A resposta será:

```json
{
    "id": 1,
    "nome": "João Silva",
    "email": "joao@email.com"
}
```

Para buscar o segundo cliente:

```text
GET http://localhost:8080/clientes/2
```

Resposta:

```json
{
    "id": 2,
    "nome": "Maria Souza",
    "email": "maria@email.com"
}
```

### 6. Testar cliente inexistente

Para testar uma busca por um ID que não existe:

```text
GET http://localhost:8080/clientes/999
```

A API retornará:

```text
Cliente não encontrado
```

com o status HTTP:

```text
404 Not Found
```

### 7. Testar JSON inválido

Também é possível testar o tratamento de dados inválidos.

Utilize:

```text
POST http://localhost:8080/clientes
```

E envie, por exemplo:

```json
{
    "nome": "João Silva"
```

Como o JSON está incompleto, a API retornará:

```text
Dados inválidos
```

com o status:

```text
400 Bad Request
```

## Códigos HTTP utilizados

| Código | Significado | Situação |
|--------|-------------|----------|
| 200 | OK | Requisição realizada com sucesso |
| 201 | Created | Cliente criado com sucesso |
| 400 | Bad Request | Dados enviados são inválidos |
| 404 | Not Found | Cliente não encontrado |

## Funcionamento da aplicação

A aplicação possui uma estrutura simples.

O tipo `Cliente` representa os dados que serão armazenados:

```go
type Cliente struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
}
```

Os clientes são armazenados em uma lista:

```go
var clientes = []Cliente{}
```

Quando um cliente é cadastrado através do método `POST`, a aplicação recebe os dados em JSON, cria um novo cliente, gera seu ID e adiciona o cliente à lista.

O método `GET /clientes` retorna todos os clientes cadastrados.

O método `GET /clientes/{id}` procura um cliente específico através do ID informado na URL.

## Armazenamento dos dados

Este projeto utiliza armazenamento **em memória**.

Isso significa que os clientes ficam armazenados apenas enquanto a aplicação estiver executando.

Ao encerrar o servidor:

```text
Ctrl + C
```

os dados serão perdidos.

Ao executar novamente:

```bash
go run .
```

a lista de clientes começará novamente vazia.

Esse comportamento é esperado para este projeto, pois o objetivo é demonstrar a criação e utilização de uma API REST em Go sem utilizar um banco de dados.

## Resumo

A API permite realizar as seguintes operações:

- Cadastrar clientes;
- Listar todos os clientes;
- Buscar clientes pelo ID;
- Validar dados enviados;
- Retornar códigos de status HTTP apropriados;
- Trabalhar com requisições e respostas em JSON.

## Comandos principais

Instalar dependência:

```bash
go get -u github.com/gorilla/mux
```

Organizar dependências:

```bash
go mod tidy
```

Executar o projeto:

```bash
go run .
```

Parar o servidor:

```text
Ctrl + C
```

## Autor

Desenvolvido como parte de um projeto de estudos em **Go (Golang)**.