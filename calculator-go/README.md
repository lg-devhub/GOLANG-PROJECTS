# Calculadora em Go

Projeto desenvolvido em **Go (Golang)** com o objetivo de implementar as quatro operações matemáticas básicas de uma calculadora:

- Soma
- Subtração
- Multiplicação
- Divisão

Além da implementação das funções, foram criados **testes automatizados** utilizando o pacote `testing` da linguagem Go para verificar se os resultados obtidos são os esperados.

## Tecnologias utilizadas

- Go (Golang)
- Pacote `testing`
- Testes automatizados

## Estrutura do projeto

```text
calculadora/
│
├── go.mod
├── calculadora.go
├── calculadora_test.go
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
cd calculadora
```

> **Importante:** execute os comandos a partir da pasta onde está localizado o arquivo `go.mod`.

### 3. Inicializar o módulo

Caso o projeto ainda não possua um `go.mod`, execute:

```bash
go mod init calculadora
```

Esse comando cria o arquivo `go.mod`, responsável por identificar e gerenciar o módulo do projeto.

## Funcionamento da calculadora

O arquivo `calculadora.go` contém as quatro funções responsáveis pelas operações matemáticas.

### Soma

A função `Soma` recebe dois números e retorna o resultado da adição:

```go
func Soma(a, b float64) float64 {
    return a + b
}
```

Exemplo:

```text
Soma(10, 5) → 15
```

### Subtração

A função `Subtracao` realiza a subtração entre dois números:

```go
func Subtracao(a, b float64) float64 {
    return a - b
}
```

Exemplo:

```text
Subtracao(10, 5) → 5
```

### Multiplicação

A função `Multiplicacao` realiza a multiplicação entre dois números:

```go
func Multiplicacao(a, b float64) float64 {
    return a * b
}
```

Exemplo:

```text
Multiplicacao(10, 5) → 50
```

### Divisão

A função `Divisao` realiza a divisão entre dois números:

```go
func Divisao(a, b float64) float64 {
    return a / b
}
```

Exemplo:

```text
Divisao(10, 5) → 2
```

## Testes automatizados

Para verificar se as funções estão funcionando corretamente, foi utilizado o pacote `testing`, disponível nativamente na linguagem Go.

Os testes estão no arquivo:

```text
calculadora_test.go
```

Cada operação possui seu próprio teste.

### Teste da soma

```go
func TestSoma(t *testing.T) {
    resultado := Soma(10, 5)

    if resultado != 15 {
        t.Errorf("Esperado 15, obtido %v", resultado)
    }
}
```

### Teste da subtração

```go
func TestSubtracao(t *testing.T) {
    resultado := Subtracao(10, 5)

    if resultado != 5 {
        t.Errorf("Esperado 5, obtido %v", resultado)
    }
}
```

### Teste da multiplicação

```go
func TestMultiplicacao(t *testing.T) {
    resultado := Multiplicacao(10, 5)

    if resultado != 50 {
        t.Errorf("Esperado 50, obtido %v", resultado)
    }
}
```

### Teste da divisão

```go
func TestDivisao(t *testing.T) {
    resultado := Divisao(10, 5)

    if resultado != 2 {
        t.Errorf("Esperado 2, obtido %v", resultado)
    }
}
```

## Executando os testes

Para executar todos os testes do projeto, utilize:

```bash
go test
```

Se todos os testes estiverem corretos, será apresentada uma saída semelhante a:

```text
PASS
ok      calculadora
```

Para visualizar detalhadamente cada teste executado, utilize:

```bash
go test -v
```

A saída será semelhante a:

```text
=== RUN   TestSoma
--- PASS: TestSoma (0.00s)
=== RUN   TestSubtracao
--- PASS: TestSubtracao (0.00s)
=== RUN   TestMultiplicacao
--- PASS: TestMultiplicacao (0.00s)
=== RUN   TestDivisao
--- PASS: TestDivisao (0.00s)
PASS
ok      calculadora
```

## Códigos dos testes

O Go utiliza algumas funções importantes do pacote `testing`.

### `testing.T`

O tipo `testing.T` permite controlar e informar o resultado de um teste.

Exemplo:

```go
func TestSoma(t *testing.T) {
    // teste
}
```

### `t.Errorf`

O método `t.Errorf` informa que o teste falhou e apresenta uma mensagem personalizada.

Exemplo:

```go
if resultado != 15 {
    t.Errorf("Esperado 15, obtido %v", resultado)
}
```

Nesse caso, se o resultado da soma for diferente de `15`, o teste será considerado como falho.

## Operações testadas

| Função | Operação | Entrada | Resultado esperado |
|--------|----------|---------|-------------------|
| `Soma` | `+` | `10, 5` | `15` |
| `Subtracao` | `-` | `10, 5` | `5` |
| `Multiplicacao` | `*` | `10, 5` | `50` |
| `Divisao` | `/` | `10, 5` | `2` |

## Conceito do desafio

O objetivo principal deste projeto é praticar:

- Criação de funções em Go;
- Parâmetros e valores de retorno;
- Utilização do tipo `float64`;
- Organização de arquivos em um projeto Go;
- Criação de testes automatizados;
- Utilização do pacote `testing`;
- Validação de resultados esperados;
- Execução de testes através do terminal.

## Comandos principais

Verificar a versão do Go:

```bash
go version
```

Inicializar o módulo:

```bash
go mod init calculadora
```

Executar os testes:

```bash
go test
```

Executar os testes mostrando detalhes:

```bash
go test -v
```

## Resultado esperado

Ao finalizar o projeto, todas as operações devem retornar os valores esperados e todos os testes devem ser aprovados.

O resultado final deverá apresentar:

```text
PASS
ok      calculadora
```

## Autor

Desenvolvido como parte de um projeto de estudos em **Go (Golang)**.