# Formatação em comandos Print

Em Go, existem diversas formas de exibir informações no console. A biblioteca padrão `fmt` oferece várias funções para formatação de saída.

## Introdução à formatação de saída

A biblioteca `fmt` (abreviação de "format") é uma parte fundamental da biblioteca padrão do Go. Ela fornece funções para formatação de entrada e saída, permitindo que você exiba informações no console de forma personalizada.

Os "verbos de formatação" disponíveis em Go são derivados da linguagem C, portanto, se você já programou em C ou linguagens similares, encontrará semelhanças na sintaxe de formatação.

## Funções Básicas de Impressão

### fmt.Println
Imprime os argumentos com espaços entre eles e uma quebra de linha no final.

```go
fmt.Println("Olá", "Mundo")  // Saída: Olá Mundo
```

### fmt.Print
Similar ao Println, mas não adiciona quebra de linha no final.

```go
fmt.Print("Olá, ")
fmt.Print("Mundo!")  // Saída: Olá, Mundo!
```

### fmt.Printf
Permite formatar a saída usando verbos de formatação.

```go
nome := "Ana"
idade := 30
fmt.Printf("%s tem %d anos\n", nome, idade)  // Saída: Ana tem 30 anos
```

Importante: Ao contrário de `fmt.Println()`, a função `fmt.Printf()` não adiciona automaticamente uma quebra de linha ao final. Se você precisar de uma quebra de linha, deve incluir `\n` no final da string de formato.

## Verbos de Formatação Comuns

### Para strings (%s)
```go
nome := "João"
fmt.Printf("Olá, %s!\n", nome)  // Saída: Olá, João!
```

### Para inteiros (%d)
```go
idade := 25
fmt.Printf("Idade: %d anos\n", idade)  // Saída: Idade: 25 anos
```

### Para números de ponto flutuante (%f)
```go
preco := 19.99
fmt.Printf("Preço: R$ %.2f\n", preco)  // Saída: Preço: R$ 19.99
```

### Para booleanos (%t)
```go
ativo := true
fmt.Printf("Status: %t\n", ativo)  // Saída: Status: true
```

### Para tipos (%T)
```go
valor := 42
fmt.Printf("O tipo de %v é %T\n", valor, valor)  // Saída: O tipo de 42 é int
```

### Para representação padrão de qualquer valor (%v)
```go
pessoa := struct {
    Nome  string
    Idade int
}{"Maria", 28}
fmt.Printf("Pessoa: %v\n", pessoa)  // Saída: Pessoa: {Maria 28}
```

### Para representação Go-syntax (%#v)
```go
pessoa := struct {
    Nome  string
    Idade int
}{"Maria", 28}
fmt.Printf("Pessoa: %#v\n", pessoa)  // Saída: Pessoa: struct { Nome string; Idade int }{Nome:"Maria", Idade:28}
```

## Formatação de Largura e Precisão

```go
// Largura mínima de 10 caracteres, alinhado à direita
fmt.Printf("%10s\n", "Olá")  // Saída: "       Olá"

// Largura mínima de 10 caracteres, alinhado à esquerda
fmt.Printf("%-10s\n", "Olá")  // Saída: "Olá       "

// Precisão de 2 casas decimais
fmt.Printf("%.2f\n", 3.14159)  // Saída: "3.14"
```

## Capturando a String Formatada

A função `Sprintf` formata e retorna uma string sem imprimi-la:

```go
nome := "Carlos"
idade := 40
mensagem := fmt.Sprintf("%s tem %d anos", nome, idade)
fmt.Println(mensagem)  // Saída: Carlos tem 40 anos
```

## Diferenças entre Println e Printf

É importante entender as diferenças entre as funções de impressão:

1. **fmt.Println**:
   - Adiciona automaticamente espaços entre os argumentos
   - Adiciona uma quebra de linha ao final
   - Não usa verbos de formatação

2. **fmt.Printf**:
   - Usa uma string de formato com verbos (%s, %d, etc.)
   - Não adiciona quebra de linha ao final (use \n se necessário)
   - Oferece controle preciso sobre a formatação

## Escolhendo a função de impressão adequada

- Use `fmt.Println` para saídas simples e rápidas
- Use `fmt.Printf` quando precisar de mais controle sobre a formatação dos valores
- Use `fmt.Print` quando precisar imprimir múltiplas partes sem quebras de linha entre elas
- Use `fmt.Sprintf` quando precisar criar uma string formatada sem imprimi-la imediatamente

## Acessando a documentação

Para uma lista completa de todos os verbos de formatação disponíveis em Go e mais detalhes sobre como usá-los, você pode consultar a documentação oficial:

```go
// No VSCode, passe o mouse sobre o pacote fmt para ver detalhes
import "fmt"
```

## Dica Prática

Para ver uma lista completa dos verbos de formatação disponíveis em Go, você pode consultar a documentação oficial da biblioteca `fmt` em: https://golang.org/pkg/fmt/ 