# Tipos compostos — arrays, slices, maps e structs

## Ao final desta aula você será capaz de...
- Diferenciar arrays de slices.
- Usar `append`, `len`, `cap` com segurança.
- Usar map para indexação rápida da agenda.

## Continuidade do mini-projeto
Evolução do Contact Book:
- `Contact` agrupa os dados de um contato;
- `map[string]Contact` usa o email como chave temporária;
- funções de listagem transformam os valores do map em `[]Contact`.

```go
type Contact struct {
    Name     string
    Email    string
    Phone    string
    Favorite bool
}

contacts := make(map[string]Contact)
contact := Contact{
    Name:     "Bia",
    Email:    "bia@example.com",
    Phone:    "+55 21 99999-9999",
    Favorite: true,
}
contacts[contact.Email] = contact
```

## Pontos da aula
- Arrays: tamanho fixo (base conceitual).
- Slices: estrutura principal do dia a dia.
- Maps: `make`, leitura, remoção, `comma ok`.

## Erros comuns que você pode encontrar
- `index out of range`.
- `assignment to entry in nil map` (esqueceu `make(map[...])`).
- Confundir nil slice com slice vazio em validações.

## Exercício prático
Refatore o Contact Book para:
1. representar contatos com `Contact`;
2. armazená-los em `map[string]Contact` usando `Contact.Email` como chave temporária;
3. implementar `findByEmail` e `removeByEmail`.

### Entrega esperada
- `list` retorna os contatos em um slice.
- Map inicializado com `make`.
- Busca por email feita diretamente no map.

### Desafio extra (opcional)
Implementar `listFavorites` retornando apenas contatos favoritos.

Continue em [Contact Book: armazenamento](contact-book.md).
