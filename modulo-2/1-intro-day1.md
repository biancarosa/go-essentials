# Seu primeiro programa em Go

Este documento apresenta os conceitos básicos para criar e executar seu primeiro programa em Go.

## Introdução

Depois de conhecer um pouco sobre a história e o contexto da linguagem Go, incluindo onde e para que foi criada, e após configurar seu ambiente de desenvolvimento com a instalação do kit de ferramentas do Go (compilador e outras utilidades) e o editor de código, chegou o momento de começar a programar de fato.

Esta etapa marca o início da nossa jornada prática com Go, onde aplicaremos os conceitos aprendidos até agora e criaremos nossos primeiros programas funcionais.

## Criando seu primeiro arquivo Go no VSCode

Para criar seu primeiro programa em Go, siga estes passos no VSCode:

1. Abra o VSCode
2. Clique em "New File" ou use o atalho Ctrl+N (Cmd+N no macOS)
3. Salve o arquivo com a extensão `.go` (por exemplo, `main.go`)

O VSCode reconhecerá automaticamente que você está trabalhando com um arquivo Go, especialmente se você já instalou a extensão do Go. Você verá algumas informações úteis na barra de status do editor:

- A linguagem do arquivo (Go)
- A versão do Go instalada em sua máquina
- Configurações de indentação (espaços/tabs)

> **Nota sobre formatação:** Em Go, você não precisa se preocupar muito com formatação do código, como espaços vs. tabs. O formatador padrão do Go (`gofmt`) que é executado automaticamente pela extensão do VSCode padroniza a formatação do seu código.

## O que é um programa Go?

Um programa Go é um arquivo de texto com a extensão `.go` que contém código escrito na linguagem de programação Go. Todo programa Go executável deve ter uma função `main()` dentro de um pacote chamado `main`.

## Estrutura básica de um programa Go

```go
// Declaração do pacote - todo programa executável deve ser do pacote main
package main

// Importação de pacotes necessários
import (
    "fmt"  // O pacote fmt contém funções para entrada/saída formatada
)

// Função principal que é executada quando o programa inicia
func main() {
    // Corpo da função - seu código vai aqui
    fmt.Println("Olá, mundo! Este é meu primeiro programa em Go!")
}
```

## Explicação dos componentes

1. **Declaração de pacote (`package main`)**: 
   Todo arquivo Go começa com uma declaração de pacote. Para programas executáveis, este pacote deve ser chamado `main`. Os pacotes em Go são como pastas - eles podem conter múltiplos arquivos, e cada arquivo no mesmo diretório pertence ao mesmo pacote.

2. **Importações (`import "fmt"`)**: 
   A seção de importação lista os pacotes que seu programa utilizará. O pacote `fmt` (formato) é frequentemente usado para funções de entrada e saída. Se você passar o mouse sobre o pacote no VSCode, verá uma descrição do pacote e poderá acessar sua documentação oficial.

3. **Função main (`func main()`)**: 
   É o ponto de entrada para seu programa. Quando você executa um programa Go, ele começa a execução pela função `main`. Esta função é especial e precisa ter exatamente este nome para programas executáveis.

4. **Comandos**: 
   Dentro da função main, você escreve comandos que seu programa executará, como exibir mensagens usando `fmt.Println()`.

É importante entender a hierarquia: pacotes contêm arquivos, arquivos contêm funções. Um diretório geralmente corresponde a um pacote, e todos os arquivos nesse diretório pertencem ao mesmo pacote.

## O programa mais simples possível

O programa Go mais simples possível tem apenas a declaração do pacote e uma função main vazia:

```go
package main

func main() {
    // Não faz nada, mas é um programa válido!
}
```

Este programa é válido, compila corretamente, mas não faz nada quando executado.

## Como executar um programa Go

Para executar seu programa no VSCode, você pode usar o terminal integrado:

1. Abra o terminal no VSCode (Menu: Terminal > New Terminal, ou use o atalho)
2. Navegue até o diretório onde seu arquivo Go está salvo (use o comando `cd`)
3. Execute o programa com o comando `go run`:

```bash
go run main.go
```

Se estiver tudo correto, seu programa será executado e você verá a saída no terminal. Se houver erros, o Go fornecerá mensagens indicando o problema.

Existem duas formas principais de executar um programa Go:

### 1. Usando `go run`

O comando `go run` compila e executa seu programa em um único passo:

```bash
go run meu_programa.go
```

Este é um método rápido para testar seu código durante o desenvolvimento.

### 2. Usando `go build` e depois executando o binário

O comando `go build` compila seu programa em um executável binário:

```bash
go build meu_programa.go
```

Isso cria um arquivo executável que você pode rodar diretamente:

```bash
# No Linux/macOS
./meu_programa

# No Windows
meu_programa.exe
```

Este método é usado quando você deseja distribuir seu programa ou executá-lo várias vezes sem recompilação.

## Verificando sua instalação Go

Para verificar se o Go está corretamente instalado e disponível no seu terminal, você pode usar o comando:

```bash
which go    # No macOS/Linux
where go    # No Windows
```

Este comando mostrará o caminho onde o Go está instalado. Se não retornar nada, pode haver um problema com sua instalação ou com as variáveis de ambiente.

## Exemplo prático

Vamos criar um programa simples que realiza uma soma e exibe o resultado:

```go
package main

import "fmt"

func main() {
    // Declarando e inicializando variáveis
    a := 5
    b := 3
    
    // Calculando a soma
    soma := a + b
    
    // Exibindo o resultado
    fmt.Println("A soma de", a, "e", b, "é", soma)
}
```

Este programa declara duas variáveis (`a` e `b`), calcula a soma delas e exibe o resultado no console.

## Próximos passos

Depois de criar e executar seu primeiro programa Go com sucesso, você pode considerar:

1. Salvar seu código em um repositório Git para acompanhar seu progresso
2. Experimentar modificações no código para entender melhor como funciona
3. Explorar outros aspectos da linguagem como:
   - Outros tipos de dados em Go
   - Estruturas de controle de fluxo (if, for, switch)
   - Funções e pacotes
   - Estruturas de dados mais complexas

Conforme avançarmos no curso, abordaremos esses tópicos de forma prática e detalhada, construindo programas cada vez mais complexos e úteis. 