# Modelo

O modelo representa os dados de um contato e é responsável por criar novos IDs.

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

O ID é criado uma vez e não deve mudar. Nome, email, telefone e favorito são atributos editáveis.

Email não é a chave primária porque pode mudar e porque sua política de unicidade pertence às regras do produto, não à identidade técnica do registro.
