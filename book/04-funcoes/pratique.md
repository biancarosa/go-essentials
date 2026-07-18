# Pratique — Funções

## Revisão rápida

1. Qual é a estrutura básica de uma função em Go?
2. Como uma função declara parâmetros?
3. Como uma função declara tipo de retorno?
4. Por que múltiplos retornos são tão comuns em Go?
5. O que significa retornar `(valor, error)`?
6. Quando uma função variádica pode ser útil?

## Leia o código

```go
func findByEmail(email string) (string, bool) {
    if email == "bia@example.com" {
        return "Bia - bia@example.com", true
    }

    return "", false
}
```

1. Quantos valores essa função retorna?
2. O que significa o `bool` nesse exemplo?
3. Qual retorno você espera para `findByEmail("unknown@example.com")`?
4. Como você chamaria essa função tratando os dois retornos?

## Pensando na agenda

1. Que assinatura você usaria para `addContact`?
2. Que assinatura você usaria para `countFavorites`?
3. Quando uma busca deveria retornar `error` em vez de apenas `bool`?
4. Por que `findContactByEmail` é mais claro que `find`?
