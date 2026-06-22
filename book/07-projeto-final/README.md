# Projeto final — Contact Book

O resultado do curso é uma pequena aplicação Go que gerencia contatos em memória.

Ela demonstra:

- tipos e valores;
- controle de fluxo;
- funções com múltiplos retornos;
- tratamento explícito de erros;
- structs, slices e maps;
- pacotes e módulos;
- dependência externa;
- UUID como identidade estável.

O projeto não inclui banco de dados, HTTP ou interface interativa. Esses elementos aumentariam o escopo sem ajudar a praticar os fundamentos.

## Critérios de conclusão

- [ ] `go run .` executa sem erro.
- [ ] `go test ./...` passa.
- [ ] Cada contato recebe um UUID.
- [ ] O store usa o ID como chave.
- [ ] É possível adicionar, buscar, listar e remover contatos.
- [ ] IDs duplicados e contatos inexistentes retornam erros.
- [ ] O código está separado em pacotes.

Use o [exemplo final](https://github.com/biancarosa/go-essentials/tree/main/examples/contact-book/06-final) apenas para comparar depois de concluir sua implementação.
