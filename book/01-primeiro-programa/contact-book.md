# Contact Book: primeiro contato

Nesta etapa, o programa apenas imprime dados fixos.

Crie um diretório e entre nele:

```bash
mkdir contact-book
cd contact-book
mkdir 01-first-program
cd 01-first-program
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
    fmt.Println("Name: Bia")
    fmt.Println("Email: bia@example.com")
    fmt.Println("Go:", runtime.Version())
}
```

Execute:

```bash
go run main.go
```

Confira o [exemplo completo da Aula 1](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/01-first-program).
