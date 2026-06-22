# Pratique — Variáveis, tipos e constantes

## Revisão rápida

1. Quando você usaria `var name string` em vez de `name := "Bia"`?
2. Por que `:=` não funciona fora de uma função?
3. O que é zero value em Go?
4. Qual é o zero value de `int`, `float64`, `string` e `bool`?
5. Por que Go exige conversão explícita entre tipos numéricos diferentes?

## Strings sem tropeço

1. Por que a frase "string em Go é sempre UTF-8" é simplificação demais?
2. O que significa dizer que string é uma sequência imutável de bytes?
3. Por que `len("ação")` pode não bater com a quantidade de caracteres que você enxerga?
4. Quando `rune` começa a fazer mais sentido do que `byte`?

## Leia o código

```go
package main

import "fmt"

func main() {
    var phone string
    favorite := true

    fmt.Printf("phone=%q favorite=%t\n", phone, favorite)
}
```

1. Qual valor aparece para `phone`?
2. Qual é o tipo inferido de `favorite`?
3. O programa compila se `favorite` não for usado? Por quê?

## Pensando na agenda

1. Quais tipos você usaria para `ID`, `Name`, `Email`, `Phone` e `Favorite`?
2. Um telefone deveria ser `int` ou `string`? Explique a escolha.
3. Por que o ID deve permanecer estável mesmo quando o email mudar?
4. Que constante faria sentido criar para o Contact Book nesta aula?
