# Go Essentials

Material aberto do curso Go Essentials, organizado como e-book e publicado com HonKit.

- O conteúdo publicado está em [`book/`](./book/).
- Os exemplos executáveis estão em [`examples/contact-book/`](./examples/contact-book/).
- A edição anterior do curso está preservada na tag [`v0`](https://github.com/biancarosa/go-essentials/tree/v0).

Todo o conteúdo explicativo está em português. Identificadores, pacotes, funções e variáveis de código usam inglês.

## Desenvolvimento local

Instale as dependências e visualize o livro:

```bash
npm install
npm run serve
```

Gere o site estático em `_book/`:

```bash
npm run build
```

## Cloudflare Pages

Use estas configurações ao conectar o repositório ao Cloudflare Pages:

- Branch de produção: `main`
- Comando de build: `npm run build`
- Diretório de saída: `_book`

Os exemplos usam módulos Go independentes. Para validar todos:

```bash
find examples/contact-book -name go.mod -execdir go test ./... \;
```
