# Funções e tratamento de erros

## Ao final desta aula você será capaz de...
- Criar funções com parâmetros e retorno.
- Aplicar múltiplos retornos no padrão Go (`valor, err`).
- Organizar regras da agenda em funções pequenas e reutilizáveis.

## Continuidade do mini-projeto
Extrair operações para funções com identificadores em inglês:
- `addContact(...)`
- `findContactByID(...)`
- `countFavorites(...)`

Exemplo de assinatura:
```go
func findContactByID(contacts []Contact, id string) (Contact, error)
```

## Pontos da aula
- Função básica e retorno tipado.
- Múltiplos retornos e tratamento explícito de erro.
- Variádicas (`...`).
- Funções anônimas como apoio em exemplos.

## Erros comuns que você pode encontrar
- Ignorar `err` sem querer.
- Retornar tipo diferente da assinatura.
- `undefined` por tentar chamar função não exportada de outro pacote (preview da próxima aula).

## Exercício prático
Implementar funções para o Contact Book:
1. Adicionar contato em um slice.
2. Buscar por ID.
3. Listar apenas nomes com retorno `[]string`.

### Entrega esperada
- Pelo menos uma função retorna `(valor, error)`.
- Em `main`, todo erro retornado é tratado com `if err != nil`.

### Desafio extra (opcional)
Criar `addContacts` como função variádica para adicionar vários contatos de uma vez.

Continue em [Contact Book: operações](contact-book.md).
