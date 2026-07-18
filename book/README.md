# Go Essentials

Aprenda os fundamentos de Go construindo, passo a passo, um **Contact Book** em memória.

Ao concluir o livro, você saberá:

- criar, compilar e executar programas Go;
- trabalhar com variáveis, tipos e constantes;
- controlar o fluxo com `if`, `for` e `switch`;
- criar funções e tratar erros explicitamente;
- organizar dados com structs, slices e maps;
- separar código em pacotes e gerenciar dependências com Go Modules;
- gerar UUIDs e usá-los como identidade estável dos contatos.

> **Nota:** Este livro não exige banco de dados, framework web ou experiência anterior com Go. Os dados permanecem em memória para que o foco fique na linguagem.

## Projeto do curso

O projeto contínuo é um Contact Book. Nas primeiras aulas, cada contato começa com dados simples:

```go
type Contact struct {
    Name     string
    Email    string
    Phone    string
    Favorite bool
}
```

Na última aula, você aprende módulos e dependências, adiciona um ID estável com `github.com/google/uuid` e passa a usar `map[uuid.UUID]Contact` como armazenamento.

Comece por [Como estudar](00-comecando/como-estudar.md).
