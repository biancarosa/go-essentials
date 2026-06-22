# Seu primeiro programa em Go

## Ao final desta aula você será capaz de...
- Instalar o Go e validar o ambiente.
- Rodar `go run main.go` e entender a estrutura mínima de um programa Go.
- Iniciar o mini-projeto da agenda com um primeiro arquivo executável.

## História do Go (jeito certo e sem confusão)
Go começou a ser desenhada dentro do Google em **2007** (Robert Griesemer, Rob Pike e Ken Thompson) e foi lançada publicamente como open source em **2009**.

## Fluxo da aula (live coding)
1. Instalação (`go version`, `go env`, VS Code + extensão Go).
2. Hello world com `package main`, `import "fmt"`, `func main()`.
3. Diferença prática entre `go run` e `go build`.
4. Projeto prático: criar `main.go` do Contact Book e imprimir um contato com UUID fixo.

Exemplo inicial:
```go
package main

import "fmt"

func main() {
    fmt.Println("Contact Book")
    fmt.Println("ID: 550e8400-e29b-41d4-a716-446655440000")
    fmt.Println("Name: Bia")
    fmt.Println("Email: bia@example.com")
}
```

## Erros comuns que você pode encontrar
- `declared and not used`
- `imported and not used`
- `undefined: fmt` (esqueceu import)
- `expected 'package', found ...` (arquivo sem `package main`)

## Exercício prático
Criar um `main.go` que imprime:
- Título: `Contact Book`
- Dois contatos com UUIDs fixos
- Versão do Go via `runtime.Version()`

### Entrega esperada
- O comando `go run main.go` executa sem erro.
- A saída mostra título, IDs, contatos e versão do Go.

### Desafio extra (opcional)
Adicionar uma mensagem final: `Go Essentials - Aula 1 concluída!`

Continue em [Contact Book: primeiro contato](contact-book.md).
