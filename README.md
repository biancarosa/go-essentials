# Go Essentials

Material aberto do curso Go Essentials, organizado como e-book e publicado com GitBook.

- O conteúdo publicado está em [`book/`](./book/).
- Os exemplos executáveis estão em [`examples/contact-book/`](./examples/contact-book/).
- A edição anterior do curso está preservada na tag [`v0`](https://github.com/biancarosa/go-essentials/tree/v0).

Todo o conteúdo explicativo está em português. Identificadores, pacotes, funções e variáveis de código usam inglês.

## Desenvolvimento local

Os exemplos usam módulos Go independentes. Para validar todos:

```bash
find examples/contact-book -name go.mod -execdir go test ./... \;
```
