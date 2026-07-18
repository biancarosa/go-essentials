# Contact Book: armazenamento

Agora podemos modelar o projeto conscientemente com uma struct e um map.

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

O email é uma chave temporária boa o suficiente para praticar `map`. Na Aula 6, o projeto ganha um ID próprio.

Implemente:

- `addContact`;
- `findContactByEmail`;
- `listContacts`;
- `removeContactByEmail`;
- `listFavorites`.

Na Aula 6, o armazenamento passa a usar `uuid.UUID` como chave.

Confira o [exemplo completo da Aula 5](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/05-composite-types).
