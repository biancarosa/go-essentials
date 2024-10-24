# github.com/biancarosa/go-essentials/go-modules

## Quando precisamos usar?

Ter projetos que são maiores do que um simples arquivo (> 1 arquivo).

## Porque usar?

- Pra que você consiga publicar o seu projeto (ex: criar um projeto que possa ser importado por outros projetos)
- Pra compilar o seu projeto de forma simples sem ter que descrever todos os arquivos ou seja, evitar ter que descrever todos os nomes no go run, ex: `go build` vs `go build main.go print.go`.