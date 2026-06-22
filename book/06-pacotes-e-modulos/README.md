# Pacotes, módulos e dependências

## Ao final desta aula você será capaz de...
- Separar código em pacotes.
- Inicializar e manter `go.mod`.
- Diferenciar corretamente `go get` e `go install pacote@versao`.
- Gerar IDs com `github.com/google/uuid`.

## Continuidade do mini-projeto
Fechamento do Contact Book:
- `main.go` (execução)
- `model/` (tipos)
- `contactbook/` (regras e armazenamento)

## Pontos da aula
- `package main` vs pacotes reutilizáveis.
- Exportação por inicial maiúscula.
- Imports internos e externos.
- `go mod init`, `go mod tidy`, `go run .`.

## Correção técnica: módulos e instalação de ferramentas
- **`go get`**: gerencia dependências do módulo atual (atualiza `go.mod`/`go.sum`).
- **`go install pacote@versao`**: preferível para instalar binários/CLI tools.

Exemplos:
```bash
# Dependência do projeto
go get github.com/google/uuid

# Instalação de ferramenta (fora da lógica da sua app)
go install golang.org/x/tools/cmd/goimports@latest
```

## Observação sobre versão no go.mod
Evite travar exemplos didáticos em uma versão específica antiga (ex.: `go 1.22`) sem contexto. Prefira orientar: “use a versão atual instalada” e mostrar uma versão apenas como exemplo.

## Erros comuns que você pode encontrar
- `missing go.mod`
- caminho de módulo/import inconsistente
- `undefined: pacote.Funcao` por tentar acessar item não exportado

## Exercício prático
Estruture o Contact Book com os pacotes `model` e `contactbook` + `main`, seguindo o [projeto final](../07-projeto-final/README.md). Troque o tipo de `Contact.ID` de `string` para `uuid.UUID`, gere IDs com `uuid.New()` e use `map[uuid.UUID]model.Contact` como store.

### Entrega esperada
- `go run .` funciona na raiz do projeto.
- `go.mod` criado e consistente com imports locais.
- `github.com/google/uuid` adicionado com `go get`.
- Contatos são buscados e removidos pelo UUID.

### Desafio extra (opcional)
Instalar `goimports` com `go install ...@latest` e aplicar no projeto.

Continue em [Contact Book: UUID e pacotes](contact-book.md).
