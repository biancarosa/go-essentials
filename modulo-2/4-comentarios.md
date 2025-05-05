# Comentários em Go

Comentários são trechos de texto no código que são ignorados pelo compilador. Eles servem para documentar o código, explicar a lógica usada e facilitar a manutenção futura.

## Introdução aos comentários

À medida que você desenvolve programas em Go, provavelmente já notou linhas que começam com `//`. Estas são comentários - texto que é ignorado pelo compilador, mas que serve para explicar seu código para outros programadores (e para você mesmo no futuro).

Os comentários são uma parte essencial da programação, especialmente quando:
- Você precisa explicar o raciocínio por trás de uma solução específica
- O código implementa algo complexo ou não-intuitivo
- Você está trabalhando em equipe e precisa comunicar suas intenções

Lembre-se: código bem escrito deve ser auto-explicativo, mas comentários bem colocados podem fornecer contexto adicional valioso.

## Tipos de Comentários

Em Go, existem dois tipos de comentários:

### 1. Comentários de Linha Única

Começam com `//` e vão até o final da linha.

```go
// Este é um comentário de linha única
fmt.Println("Hello, World!") // Este é outro comentário de linha única
```

Os comentários de linha única são úteis tanto para explicar o que acontece na linha abaixo quanto para adicionar notas ao final de uma linha de código.

### 2. Comentários de Múltiplas Linhas

Começam com `/*` e terminam com `*/`. Podem abranger várias linhas.

```go
/* Este é um comentário
   de múltiplas linhas.
   Pode abranger várias
   linhas de código. */
```

## Boas Práticas

### Documentação de Pacotes

Comentários que precedem a declaração de um pacote são considerados documentação do pacote.

```go
// Package calculator fornece funções para realizar operações matemáticas básicas.
package calculator
```

### Documentação de Funções

Comentários que precedem a declaração de funções são considerados documentação da função.

```go
// Soma retorna a soma de dois números inteiros.
// 
// Exemplo:
//
//     resultado := Soma(5, 3)
//     fmt.Println(resultado) // Imprime: 8
//
func Soma(a, b int) int {
    return a + b
}
```

### Documentação de Tipos

Comentários que precedem a declaração de tipos são considerados documentação do tipo.

```go
// Pessoa representa uma pessoa com nome e idade.
type Pessoa struct {
    Nome  string
    Idade int
}
```

### TODOs e FIXMEs

É comum usar `TODO` e `FIXME` em comentários para marcar áreas que precisam de atenção.

```go
// TODO: Implementar validação de entrada
// FIXME: Corrigir bug quando valor é negativo
```

## Onde termina um comentário?

Para comentários de linha única (`//`), o comentário termina automaticamente quando a linha termina. Se você iniciar uma nova linha, precisará adicionar outro `//` se quiser continuar o comentário.

Para comentários de múltiplas linhas (`/* ... */`), o comentário só termina quando você fecha com `*/`, não importa quantas linhas estejam entre os delimitadores.

## Comentários e o compilador Go

É importante lembrar que os comentários são completamente ignorados pelo compilador Go. Eles estão lá apenas para os seres humanos que leem o código. Isso significa que você pode adicionar comentários livremente sem afetar o desempenho ou o comportamento do seu programa.

## Documentação com godoc

Go possui uma ferramenta chamada `godoc` que gera documentação a partir dos comentários. Esta documentação segue um formato específico:

```go
// NomeDaFuncao faz algo específico com os parâmetros fornecidos.
// Ele retorna um resultado baseado em alguma lógica.
//
// O primeiro parágrafo é um resumo.
//
// Parágrafos seguintes fornecem detalhes.
//
// Exemplos de uso podem ser mostrados usando formatação de código
// indentado com pelo menos um espaço:
//
//     resposta := NomeDaFuncao(param1, param2)
//
// Parâmetros adicionais podem ser documentados assim:
//
// nome: descrição do parâmetro
// outro: descrição de outro parâmetro
func NomeDaFuncao(nome string, outro int) string {
    // implementação
}
```

## Comentários Explicativos vs. Descritivos

- **Comentários Explicativos**: Explicam *por que* algo está sendo feito (a lógica/raciocínio)
- **Comentários Descritivos**: Descrevem *o que* está sendo feito (geralmente redundante)

Prefira comentários explicativos. Se seu código precisa de comentários descritivos, considere torná-lo mais claro.

```go
// Mal exemplo (descritivo):
// Incrementa contador
contador++

// Bom exemplo (explicativo):
// Incrementa o contador para compensar o elemento cabeçalho que não deve ser incluído no cálculo
contador++
```

## Quando Não Usar Comentários

Nem sempre você precisa de comentários. Código bem escrito, com nomes de variáveis e funções descritivos, muitas vezes é autoexplicativo.

```go
// Não necessário:
// Calcula a média
media := soma / float64(quantidade)

// Melhor deixar o código falar por si:
mediaTemperaturas := somaTemperaturas / float64(quantidadeAmostras)
```

## Uso de comentários no fluxo de trabalho com Git

Comentários também são úteis quando você está trabalhando com controle de versão. Ao fazer commits no Git, você pode usar comentários para lembrar de alterações que ainda não foram implementadas ou problemas que precisam ser resolvidos.

Quando seu código estiver pronto para ser compartilhado, você pode fazer commit e push das mudanças para seu repositório Git, incluindo os comentários que explicam seu código.

## Comandos go doc

Para visualizar a documentação de um pacote, tipo ou função, você pode usar o comando:

```bash
go doc fmt.Println    # Documentação da função Println do pacote fmt
go doc fmt            # Documentação do pacote fmt
``` 