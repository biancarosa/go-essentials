# Contact Book: UUID e pacotes

Inicialize o módulo:

```bash
go mod init github.com/biancarosa/contact-book
go get github.com/google/uuid
```

O projeto final terá esta estrutura:

```text
contact-book/
├── go.mod
├── main.go
├── model/
│   └── contact.go
└── contactbook/
    └── store.go
```

No pacote `model`, o ID passa a ser um UUID real:

```go
package model

import "github.com/google/uuid"

type Contact struct {
    ID       uuid.UUID
    Name     string
    Email    string
    Phone    string
    Favorite bool
}

func NewContact(name, email, phone string, favorite bool) Contact {
    return Contact{
        ID:       uuid.New(),
        Name:     name,
        Email:    email,
        Phone:    phone,
        Favorite: favorite,
    }
}
```

O store usa exatamente o mesmo tipo como chave:

```go
type Store map[uuid.UUID]model.Contact
```

Confira o [projeto final executável](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/06-final).
