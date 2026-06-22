# Preparando o ambiente

Você precisa do Go, de um editor e de um terminal.

## macOS

Baixe o instalador em [go.dev/dl](https://go.dev/dl/) ou use Homebrew:

```bash
brew install go
```

## Windows

Baixe o instalador em [go.dev/dl](https://go.dev/dl/) e siga o assistente. Depois, abra um novo PowerShell.

## Linux

Siga as instruções da [documentação oficial](https://go.dev/doc/install) para instalar a versão atual.

Confirme a instalação:

```bash
go version
go env
```

## Editor

Recomendamos o Visual Studio Code com a extensão oficial **Go**, publicada pela Go Team at Google. A extensão oferece formatação, autocomplete, análise estática e integração com o debugger.

## Checklist

- [ ] `go version` mostra a versão instalada.
- [ ] `go env` executa sem erro.
- [ ] O editor reconhece arquivos `.go`.
- [ ] `gofmt` está disponível.
- [ ] O terminal consegue executar `go run`.

Com o ambiente pronto, avance para [Seu primeiro programa em Go](../01-primeiro-programa/README.md).
