# compiling-go-code

1 - Go é uma linguagem COMPILADA - muito parecido com C!

CÓDIGO -> COMPILADOR DO CÓDIGO -> LINGUAGEM DE MÁQUINA -> LINKING (LINKAGEM) -> EXECUTÁVEL (PROGRAMA)

main.go -> compilador do Go -> executável main

- Go é uma linguagem multi-plataforma! 

- Linguagens de programação podem ser interpretadas ou compiladas!

- (não é sempre) linguagens de programação compiladas tendem a ser mais performáticas

- `go run main.go` -> compila e roda o arquivo main.go
- `go build` ?


## Perguntas

- Ele compila apenas a nível de os ou de arquitetura também? Exemplo compilei em um Linux x86 eu consigo executar em um Linux arm?

- OS: Window, Linux, Mac
- Arquitetura: 64 bits, 32 bits, ARM
--- GOOS
--- GOARCH

go version go1.22.6 darwin/amd64

GOOS=windows GOARCH=amd64 go build

# Caso de uso muito útil!

- Compartilhar um binário com alguém: main num pendrive! 
- CI/CD em Go: 
------ CI: compilar
------ CD: colocar o binario pra rodar na maquina
- Máquina que vai rodar o código não precisa ter acesso a nada do código fonte!

# Porque a forma que Go compila o código é "melhor" do que em C?

Pros nossos casos de uso:
--- Binário Go é autocontido, ou seja, qualquer biblioteca de sistema etc já vem dentro do binário
--- Código compilado em C nem sempre tem dentro desse código as biblioteca (.h) que o seu código precisa pra rodar dentro do executável

# Flags de compilação

- Ajudam a customizar o binário (executável) gerado
- E as vezes ajudam a diminuir o tamanho desse executável 
- Vcs vão ver Dockerfile customizando flags pra gerar um binário menor pra rodar em ambientes produtivos