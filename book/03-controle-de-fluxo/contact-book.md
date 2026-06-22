# Contact Book: filtros

Agora usaremos `for`, `if` e `switch` para ler uma coleção de contatos fornecida pelo exercício.

```go
contacts := []string{"Bia", "Ana", "Joao"}

for _, name := range contacts {
    if name == "Ana" {
        fmt.Println(name, "is a favorite contact")
        continue
    }
    fmt.Println(name)
}

switch len(contacts) {
case 0:
    fmt.Println("empty")
case 1, 2, 3, 4, 5:
    fmt.Println("starting")
default:
    fmt.Println("growing")
}
```

O objetivo ainda não é modelar o contato perfeitamente. Estamos praticando como percorrer valores e tomar decisões.

Desafio: mantenha um segundo slice com os favoritos e conte quantos contatos possuem telefone vazio.

Confira o [exemplo completo da Aula 3](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/03-flow-control).
