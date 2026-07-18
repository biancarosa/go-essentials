# Variáveis, tipos e constantes

## Ao final desta aula você será capaz de...
- Declarar variáveis com `var` e `:=`.
- Entender zero values e conversões explícitas.
- Usar tipos básicos sem cair nas pegadinhas de string.

## Continuidade do mini-projeto
Agora o Contact Book passa a representar cada dado com uma variável. O ID fica para depois, quando a gente introduzir módulos e a dependência de UUID.

```go
contactName := "Bia"
contactEmail := "bia@example.com"
contactPhone := "+55 21 99999-9999"
contactFavorite := true
```

## Ajuste técnico importante sobre strings
Em Go, **string é uma sequência imutável de bytes**.
- Literais de string no código-fonte (ex.: `"Olá"`) são UTF-8.
- Mas uma string pode conter bytes arbitrários.

Ou seja: evite a frase “string é sempre UTF-8” sem contexto.

## Pontos da aula
- Tipos: `int`, `float64`, `bool`, `string`, `byte`, `rune`.
- Zero values.
- Conversão explícita (`int(x)`, `float64(y)`) e `strconv`.
- Constantes com `const` e `iota`.

## Erros comuns que você pode encontrar
- `cannot use X (type ...) as type ...`
- `invalid operation: mismatched types`
- `declared and not used`

## Exercício prático
Represente 3 contatos em memória com valores hardcoded e identificadores em inglês.
Mostre:
- Quantidade total de contatos
- Quantos são favoritos

### Entrega esperada
- Código compila sem warnings/erros.
- Usa pelo menos 1 constante e 1 contador.

### Desafio extra (opcional)
Mostrar no terminal um aviso quando `contactPhone` estiver vazio.

Continue em [Contact Book: dados do contato](contact-book.md).
