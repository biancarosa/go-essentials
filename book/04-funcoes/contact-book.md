# Contact Book: operações

As regras do projeto começam a sair de `main` e ganhar nomes próprios.

```go
type Contact struct {
    ID       string
    Name     string
    Email    string
    Phone    string
    Favorite bool
}

func findContactByID(contacts []Contact, id string) (Contact, error) {
    for _, contact := range contacts {
        if contact.ID == id {
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

{% hint style="info" %}
A struct é fornecida como apoio nesta etapa. A Aula 5 explica como structs e slices funcionam.
{% endhint %}

Implemente também `addContact` e trate todo erro retornado por `findContactByID`.

Confira o [exemplo completo da Aula 4](https://github.com/biancarosa/go-essentials/tree/master/examples/contact-book/04-functions).
