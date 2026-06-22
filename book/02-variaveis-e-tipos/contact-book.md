# Contact Book: dados do contato

Substitua parte da saída fixa por variáveis. Use nomes em inglês no código:

```go
package main

import "fmt"

func main() {
    const appName = "Contact Book"

    contactID := "550e8400-e29b-41d4-a716-446655440000"
    contactName := "Bia"
    contactEmail := "bia@example.com"
    contactPhone := "+55 21 99999-9999"
    contactFavorite := true

    fmt.Println(appName)
    fmt.Println("ID:", contactID)
    fmt.Println("Name:", contactName)
    fmt.Println("Email:", contactEmail)
    fmt.Println("Phone:", contactPhone)
    fmt.Println("Favorite:", contactFavorite)
}
```

O ID continua sendo uma `string` com formato de UUID. Isso nos permite trabalhar com o valor antes de aprender dependências externas.

{% hint style="warning" %}
Telefone deve ser `string`, não `int`. Símbolos, espaços, zeros à esquerda e códigos de país fazem parte do valor.
{% endhint %}

Confira o [exemplo completo da Aula 2](https://github.com/biancarosa/go-essentials/tree/master/examples/contact-book/02-variables-and-types).
