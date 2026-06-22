# Controle de fluxo — if, for e switch

## Ao final desta aula você será capaz de...
- Escrever decisões com `if/else`.
- Repetir com `for` (clássico, estilo while e `range`).
- Usar `switch` para regras mais legíveis.

## Continuidade do mini-projeto
Vamos aplicar regras no Contact Book:
- listar contatos favoritos;
- buscar por prefixo de nome;
- classificar contatos por status (`"new"`, `"active"`, etc.) com `switch`.

## Pontos da aula
- `if` sem parênteses, com short statement.
- `for` como único loop do Go.
- `break` e `continue`.
- `switch` com e sem expressão.

## Erros comuns que você pode encontrar
- Loop infinito sem condição de saída.
- `index out of range` ao percorrer slice com limite errado.
- Esquecer `continue`/`break` e ter comportamento inesperado.

## Exercício prático
Com base em um slice de contatos fornecido pela aula:
1. Imprima todos os campos `Name`.
2. Imprima somente contatos com `Favorite == true`.
3. Conte quantos têm `Phone == ""`.

### Entrega esperada
- Usa ao menos 1 `if`, 1 `for range` e 1 `switch`.
- Não gera `panic` em execução.

### Desafio extra (opcional)
Criar uma função que retorna um status pela quantidade de contatos:
- 0: `"empty"`
- 1–5: `"starting"`
- 6+: `"growing"`

Continue em [Contact Book: filtros](contact-book.md).
