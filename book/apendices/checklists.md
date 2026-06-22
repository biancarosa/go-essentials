# Checklists

## Antes de começar

- [ ] Go instalado e disponível em `go version`.
- [ ] Editor reconhece arquivos `.go`.
- [ ] Extensão oficial do Go instalada.
- [ ] Terminal aberto no diretório correto.

## Antes de entregar um exercício

- [ ] Execute `go fmt ./...`.
- [ ] Execute `go test ./...` quando houver testes.
- [ ] Execute `go run .`.
- [ ] Remova código e imports não usados.
- [ ] Confira os critérios de aceitação.
- [ ] Use identificadores de código em inglês.

## Projeto final

- [ ] `Contact.ID` usa `uuid.UUID`.
- [ ] O store usa o ID como chave.
- [ ] Email é apenas atributo e busca secundária.
- [ ] Erros são verificados.
- [ ] O código está dividido em `model`, `contactbook` e `main`.
