# Contact Book: dados do contato

Substitua parte da saída fixa por variáveis. Use nomes em inglês no código:

```go
package main

import "fmt"

func main() {
    const appName = "Contact Book"

    contactName := "Bia"
    contactEmail := "bia@example.com"
    contactPhone := "+55 21 99999-9999"
    contactFavorite := true

    fmt.Println(appName)
    fmt.Println("Name:", contactName)
    fmt.Println("Email:", contactEmail)
    fmt.Println("Phone:", contactPhone)
    fmt.Println("Favorite:", contactFavorite)
}
```

O ID entra mais tarde, quando a gente introduzir módulos e a dependência que gera UUIDs.

> **Atenção:** Telefone deve ser `string`, não `int`. Símbolos, espaços, zeros à esquerda e códigos de país fazem parte do valor.

Confira o [exemplo completo da Aula 2](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/02-variables-and-types).
