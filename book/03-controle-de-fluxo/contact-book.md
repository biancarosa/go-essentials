# Contact Book: filtros

Agora usaremos `for`, `if` e `switch` para ler uma coleção de contatos fornecida pelo exercício.

```go
package main

import "fmt"

type Contact struct {
    Name     string
    Phone    string
    Favorite bool
}

func main() {
    contacts := []Contact{
        {Name: "Bia", Phone: "+55 21 99999-9999", Favorite: true},
        {Name: "Ana", Phone: "", Favorite: false},
        {Name: "Joao", Phone: "+55 11 99999-9999", Favorite: true},
    }

    emptyPhones := 0

    for _, contact := range contacts {
        if contact.Favorite {
            fmt.Println("Favorite:", contact.Name)
        }

        if contact.Phone == "" {
            emptyPhones++
        }
    }

    fmt.Println("Contacts without phone:", emptyPhones)

    switch len(contacts) {
    case 0:
        fmt.Println("empty")
    case 1, 2, 3, 4, 5:
        fmt.Println("starting")
    default:
        fmt.Println("growing")
    }
}
```

O objetivo ainda não é modelar o contato perfeitamente. A struct está pronta para que você possa focar em percorrer valores e tomar decisões.

Desafio: imprima todos os nomes, depois ajuste os dados para cair nos outros ramos do `switch`.

Confira o [exemplo completo da Aula 3](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/03-flow-control).
