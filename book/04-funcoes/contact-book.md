# Contact Book: operações

As regras do projeto começam a sair de `main` e ganhar nomes próprios.

```go
type Contact struct {
    Name     string
    Email    string
    Phone    string
    Favorite bool
}

func findContactByEmail(contacts []Contact, email string) (Contact, error) {
    for _, contact := range contacts {
        if contact.Email == email {
            return contact, nil
        }
    }
    return Contact{}, fmt.Errorf("contact not found")
}

func countFavorites(contacts []Contact) int {
    count := 0
    for _, contact := range contacts {
        if contact.Favorite {
            count++
        }
    }
    return count
}
```

> **Nota:** A struct é fornecida como apoio nesta etapa. A Aula 5 explica como structs e slices funcionam.

Implemente também `addContact` e trate todo erro retornado por `findContactByEmail`.

Confira o [exemplo completo da Aula 4](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/04-functions).
