# Pratique — Seu primeiro programa

## Revisão rápida

1. O que precisa existir em um arquivo Go para ele virar um programa executável?
2. Qual é a diferença prática entre `go run main.go` e `go build`?
3. Para que serve a linha `package main`?
4. Por que `func main()` é especial em um programa Go?
5. O que o pacote `fmt` faz no primeiro exemplo da aula?

## Leia o código

```go
package main

import "fmt"

func main() {
    message := "Contact Book"
    fmt.Println(message)
}
```

1. Qual é a saída esperada desse programa?
2. O que acontece se removermos a linha `fmt.Println(message)`?
3. O que acontece se removermos o `import "fmt"` mas mantivermos `fmt.Println`?

## Pensando na agenda

1. Que informações mínimas, incluindo o UUID, um contato precisa ter para aparecer no primeiro programa?
2. Por que começar com contatos fixos no código é suficiente nesta aula?
3. Como você mudaria a saída para mostrar dois contatos em vez de um?
