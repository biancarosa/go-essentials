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

O projeto contínuo é um Contact Book. Cada contato possui:

```go
type Contact struct {
    ID       uuid.UUID
    Name     string
    Email    string
    Phone    string
    Favorite bool
}
```

Nas primeiras aulas, os UUIDs são fornecidos como strings fixas. Na última aula, você aprende módulos e dependências, passa a gerar IDs com `github.com/google/uuid` e usa `map[uuid.UUID]Contact` como armazenamento.

Comece por [Como estudar](00-comecando/como-estudar.md).
