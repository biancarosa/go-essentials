# Pratique — Pacotes e módulos

## Revisão rápida

1. Qual é a diferença entre `package main` e um pacote reutilizável?
2. Como Go decide se um nome é exportado por um pacote?
3. Por que a convenção é o nome do pacote bater com o nome do diretório?
4. Para que serve `go mod init`?
5. Que informações básicas aparecem em um `go.mod`?
6. Para que serve o `go.sum`?

## Dependências e ferramentas

1. Quando usar `go get`?
2. Quando usar `go install pacote@versao`?
3. Por que instalar uma CLI com `go get` não é mais o caminho preferido?
4. O que `go mod tidy` faz?
5. Quando `go mod download` pode ser útil?

## Leia a estrutura

```text
contact-book/
├── go.mod
├── main.go
├── model/
│   └── contact.go
└── contactbook/
    └── store.go
```

1. Que pacote você espera encontrar em `main.go`?
2. Que pacote você espera encontrar em `contactbook/store.go`?
3. Como `main.go` poderia importar `model` e `contactbook` se o módulo fosse `github.com/biancarosa/contact-book`?
4. O que acontece se uma função em `contactbook` começar com letra minúscula e você tentar chamá-la de `main`?

## Pensando na agenda

1. Que código deveria ficar no pacote `model`?
2. Que código deveria ficar no pacote `contactbook`?
3. Que código deveria continuar no `main`?
4. Por que o store final usa `map[uuid.UUID]model.Contact`?
5. Por que `github.com/google/uuid` é uma dependência do projeto e não uma ferramenta instalada com `go install`?
