# Pratique — Controle de fluxo

## Revisão rápida

1. Por que o `if` em Go não usa parênteses obrigatórios?
2. O que é um `if` com short statement?
3. Quais são as três formas principais de usar `for` em Go?
4. Como `for range` facilita percorrer slices, maps e strings?
5. Qual é a diferença entre `break` e `continue`?
6. Por que `switch` em Go não precisa de `break` em cada `case`?

## Leia o código

```go
contacts := []string{"Bia", "Ana", "Joao"}

for i, name := range contacts {
    if name == "Ana" {
        continue
    }

    fmt.Println(i, name)
}
```

1. Quais nomes são impressos?
2. O índice de `"Joao"` muda por causa do `continue`?
3. O que mudaria se fosse usado `break` no lugar de `continue`?

## Pensando na agenda

1. Como você listaria apenas contatos com `Favorite == true`?
2. Que regra com `if/else` poderia detectar `Phone == ""`?
3. Que classificação com `switch` faria sentido para uma agenda simples?
4. Como evitar `index out of range` ao acessar um contato por posição?
5. Por que buscar por UUID em um map não depende da posição do contato?
