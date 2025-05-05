# Valores que podemos ter em Go

Go é uma linguagem estaticamente tipada, o que significa que cada variável possui um tipo específico que determina quais valores ela pode armazenar. Aqui estão os principais tipos de dados em Go:

## Introdução aos tipos de valores

Quando começamos a programar em Go, é essencial entender os diferentes tipos de valores que podemos usar. Estes tipos são os blocos fundamentais para construir nossos programas.

Em nosso primeiro programa, já utilizamos valores como strings (cadeias de caracteres), mas Go oferece muitos outros tipos que podemos explorar. Vamos conhecer os principais tipos de valores que você usará com frequência:

- **Strings**: Sequências de caracteres, como "Olá, mundo!"
- **Números**: Inteiros e números de ponto flutuante
- **Expressões**: Combinações de valores, como 2 + 5
- **Booleanos**: Valores verdadeiro (`true`) ou falso (`false`)

Conforme você avança na linguagem, você descobrirá como trabalhar com esses diferentes tipos de valores para resolver problemas cada vez mais complexos.

## Tipos Básicos

### Números Inteiros
- **int**: Inteiro com sinal, tamanho depende da arquitetura (32 ou 64 bits)
- **int8**, **int16**, **int32**, **int64**: Inteiros com sinal de tamanho específico
- **uint**: Inteiro sem sinal, tamanho depende da arquitetura
- **uint8**, **uint16**, **uint32**, **uint64**: Inteiros sem sinal de tamanho específico
- **byte**: Alias para uint8, usado para representar bytes
- **rune**: Alias para int32, usado para representar caracteres Unicode

### Números de Ponto Flutuante
- **float32**: Ponto flutuante de precisão simples (32 bits)
- **float64**: Ponto flutuante de precisão dupla (64 bits)

### Números Complexos
- **complex64**: Números complexos com partes real e imaginária float32
- **complex128**: Números complexos com partes real e imaginária float64

### Booleanos
- **bool**: Valores verdadeiro (`true`) ou falso (`false`)

### Texto
- **string**: Sequência imutável de bytes, geralmente representando texto UTF-8

## Tipos Compostos

### Arrays
- Coleções de tamanho fixo de elementos do mesmo tipo
- Exemplo: `[5]int` é um array de 5 inteiros

### Slices
- Referências para seções de arrays, com tamanho dinâmico
- Exemplo: `[]int` é um slice de inteiros

### Maps
- Coleções não-ordenadas de pares chave-valor
- Exemplo: `map[string]int` é um map com chaves string e valores inteiros

### Structs
- Tipos compostos que agrupam campos de dados
- Exemplo: `type Pessoa struct { Nome string; Idade int }`

## Declaração de Variáveis

Em Go, você pode declarar variáveis de várias maneiras:

```go
// Declaração explícita
var idade int = 25

// Inferência de tipo
var nome = "Maria"

// Declaração curta (apenas dentro de funções)
temperatura := 23.5
```

## Utilizando diferentes tipos em expressões

Em Go, você pode combinar diferentes tipos de valores em expressões, desde que sejam compatíveis:

```go
// Expressão com números
soma := 2 + 5 // Resultado: 7

// Expressão com strings
saudacao := "Olá, " + "mundo!" // Resultado: "Olá, mundo!"

// Usando valores booleanos em expressões
resultado := true && false // Resultado: false
```

Estas expressões podem ser usadas diretamente em funções como `fmt.Println()`:

```go
fmt.Println(2 + 5) // Imprime: 7
fmt.Println(true || false) // Imprime: true
```

## Zero Values

Quando você declara uma variável sem atribuir um valor, Go atribui automaticamente o "valor zero" para o tipo:

- Números: `0`
- Booleanos: `false`
- Strings: `""`
- Ponteiros, funções, interfaces, slices, channels, maps: `nil` 