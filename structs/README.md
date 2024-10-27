
# Structs

## O que são structs em Go?

No Go, uma struct (ou estrutura) é um tipo de dado composto que agrupa campos (também chamados de atributos ou propriedades) sob um único nome. Cada campo tem um nome e um tipo. As structs são semelhantes às classes em outras linguagens de programação orientadas a objetos, mas o Go não tem classes nem herança. Em vez disso, usa structs e interfaces para organizar e reutilizar o código.

## Por que usar structs?
As structs são úteis quando você quer representar algo que tem múltiplos atributos relacionados. Elas permitem organizar melhor os dados, facilitando a leitura e manutenção do código. Além disso, ao usar structs, você pode:

- Agrupar dados relacionados sob um nome comum.
- Facilitar a passagem e manipulação de dados em funções.
- Usar métodos associados à struct para encapsular comportamentos.

## Como declarar uma struct?
Aqui está um exemplo simples de como declarar e usar uma struct em Go:

```
package main

import "fmt"

// Declaração da struct
type Pessoa struct {
    Nome   string
    Idade  int
    Cidade string
}

func main() {
    // Inicializando uma struct diretamente
    pessoa1 := Pessoa{
        Nome:   "nome",
        Idade:  29,
        Cidade: "cidade",
    }

    // Inicializando uma struct sem nomear os campos
    pessoa2 := Pessoa{"João", 34, "São Paulo"}

    // Acessando os campos da struct
    fmt.Println(pessoa1.Nome)  // Saída: nome
    fmt.Println(pessoa2.Cidade) // Saída: São Paulo

    // Atualizando um campo
    pessoa1.Idade = 30
    fmt.Printf("%s agora tem %d anos.\n", pessoa1.Nome, pessoa1.Idade)
}
```

### Componentes da struct:

- Declaração: Você usa a palavra-chave type para definir uma nova struct e dar um nome a ela. Em seguida, define os campos que pertencem a essa struct, com seus respectivos tipos. No exemplo acima, temos a struct Pessoa com três campos: Nome, Idade e Cidade.

- Instanciação: Para usar uma struct, você precisa inicializá-la. No exemplo, pessoa1 e pessoa2 são instâncias da struct Pessoa. Você pode inicializar uma struct especificando os valores de seus campos na ordem correta ou explicitamente nomeando os campos.

- Acesso e Modificação: Os campos de uma struct podem ser acessados e modificados usando a notação . (ponto). No exemplo, pessoa1.Nome acessa o campo Nome da struct pessoa1.

- Métodos em structs Embora Go não tenha classes, você pode associar métodos a uma struct. Um método é basicamente uma função que "pertence" a uma struct e tem acesso aos seus campos.

```
// Definindo um método para a struct Pessoa
func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome + "!"
}

func main() {
    pessoa := Pessoa{Nome: "Bianca", Idade: 29, Cidade: "Rio de Janeiro"}

    // Chamando o método
    fmt.Println(pessoa.Saudacao())  // Saída: Olá, meu nome é Bianca!
}
```


### Structs como valores e referências
Por padrão, quando você passa uma struct para uma função ou método, ela é passada por valor, ou seja, uma cópia da struct é feita. Se você quiser modificar a struct original, deve passar um ponteiro para ela:

```
// Método que altera o campo da struct usando ponteiro
func (p *Pessoa) Envelhecer() {
    p.Idade += 1
}

func main() {
    pessoa := Pessoa{Nome: "Marcela", Idade: 29, Cidade: "Rio de Janeiro"}

    // Chamando o método que altera a idade
    pessoa.Envelhecer()
    fmt.Printf("%s agora tem %d anos.\n", pessoa.Nome, pessoa.Idade)  // Saída: Marcela agora tem 30 anos.
}
```

Aqui, o método Envelhecer() usa um ponteiro *Pessoa, então ele modifica a instância original de Pessoa em vez de trabalhar com uma cópia.

### Comparação com classes em outras linguagens

- Go não tem herança: Ao contrário de linguagens como Java ou Python, Go não suporta herança. Em vez disso, você pode compor structs ou usar interfaces para compartilhar comportamento entre tipos diferentes.

- Não há construtores explícitos: No Go, você não precisa definir explicitamente um construtor como em outras linguagens orientadas a objetos. Em vez disso, você pode criar funções que retornam uma nova instância de uma struct, como uma forma de construtor.

#### Exemplo de "construtor"

```
func NovaPessoa(nome string, idade int, cidade string) Pessoa {
    return Pessoa{
        Nome:   nome,
        Idade:  idade,
        Cidade: cidade,
    }
}

func main() {
    pessoa := NovaPessoa("José", 29, "Recife")
    fmt.Println(pessoa)
}
```

# Links úteis
- [A Tour of Go: Structs](https://go.dev/tour/moretypes/2)
- [Go by Example: Structs](https://gobyexample.com/structs)
