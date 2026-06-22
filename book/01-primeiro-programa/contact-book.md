# Contact Book: primeiro contato

Nesta etapa, o programa apenas imprime dados fixos. Ainda não precisamos de variáveis ou estruturas de dados.

Crie um diretório e entre nele:

```bash
mkdir contact-book
cd contact-book
```

Crie `main.go`:

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Contact Book")
    fmt.Println("ID: 550e8400-e29b-41d4-a716-446655440000")
    fmt.Println("Name: Bia")
    fmt.Println("Email: bia@example.com")
    fmt.Println("Go:", runtime.Version())
}
```

Execute:

```bash
go run main.go
```

O UUID é fixo por enquanto. Ele funciona apenas como um valor que identifica o contato; aprenderemos a gerá-lo na Aula 6.

Confira o [exemplo completo da Aula 1](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/01-first-program).
