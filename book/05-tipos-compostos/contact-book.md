# Contact Book: armazenamento

Agora podemos modelar o projeto conscientemente com uma struct e um map.

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

O `ID` é a chave primária lógica. Email e telefone podem mudar sem alterar a identidade do contato.

Implemente:

- `addContact`;
- `findContactByID`;
- `listContacts`;
- `removeContactByID`;
- `findContactByEmail` como busca secundária.

Na Aula 6, `ID` deixa de ser `string` e passa a ser `uuid.UUID`.

Confira o [exemplo completo da Aula 5](https://github.com/biancarosa/go-essentials/tree/master/examples/contact-book/05-composite-types).
