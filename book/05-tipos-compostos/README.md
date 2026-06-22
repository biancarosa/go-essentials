# Tipos compostos — arrays, slices, maps e structs

## Ao final desta aula você será capaz de...
- Diferenciar arrays de slices.
- Usar `append`, `len`, `cap` com segurança.
- Usar map para indexação rápida da agenda.

## Continuidade do mini-projeto
Evolução do Contact Book:
- `Contact` agrupa os dados de um contato;
- `map[string]Contact` usa o UUID textual como chave primária lógica;
- funções de listagem transformam os valores do map em `[]Contact`.

```go
type Contact struct {
    ID       string
    Name     string
    Email    string
    Phone    string
    Favorite bool
}

contacts := make(map[string]Contact)
contact := Contact{
    ID:       "550e8400-e29b-41d4-a716-446655440000",
    Name:     "Bia",
    Email:    "bia@example.com",
    Phone:    "+55 21 99999-9999",
    Favorite: true,
}
contacts[contact.ID] = contact
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
2. armazená-los em `map[string]Contact` usando `Contact.ID` como chave;
3. implementar `findByID` e `removeByID`.

### Entrega esperada
- `list` retorna os contatos em um slice.
- Map inicializado com `make`.
- Busca por ID feita diretamente no map.

### Desafio extra (opcional)
Implementar `findByEmail` como busca secundária, sem transformar email em chave primária.

Continue em [Contact Book: armazenamento](contact-book.md).
