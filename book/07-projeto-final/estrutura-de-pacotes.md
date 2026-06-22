# Estrutura de pacotes

```text
contact-book/
├── go.mod
├── go.sum
├── main.go
├── model/
│   └── contact.go
└── contactbook/
    ├── store.go
    └── store_test.go
```

## Responsabilidades

`model` contém o tipo `Contact` e sua criação.

`contactbook` contém o armazenamento e as operações sobre contatos.

`main` cria os valores, chama os pacotes e apresenta o resultado. Ele não deve concentrar regras do Contact Book.

## Dependência

```bash
go get github.com/google/uuid
go mod tidy
```

O arquivo `go.sum` é gerenciado pelo Go e deve ser versionado junto com `go.mod`.
