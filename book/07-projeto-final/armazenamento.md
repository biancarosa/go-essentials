# Armazenamento

O store é um map em memória:

```go
type Store map[uuid.UUID]model.Contact
```

As operações públicas são:

```go
func NewStore() Store
func Add(store Store, contact model.Contact) error
func FindByID(store Store, id uuid.UUID) (model.Contact, error)
func FindByEmail(store Store, email string) (model.Contact, error)
func List(store Store) []model.Contact
func ListFavorites(store Store) []model.Contact
func Remove(store Store, id uuid.UUID) error
```

Dois erros fazem parte do contrato:

```go
var (
    ErrContactNotFound      = errors.New("contact not found")
    ErrContactAlreadyExists = errors.New("contact already exists")
)
```

> **Atenção:** A iteração de maps não possui ordem garantida. `List` não deve prometer uma ordem específica sem ordenar explicitamente o resultado.
