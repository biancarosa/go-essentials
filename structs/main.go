package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Cidade    string
}

func main() {
	var p Pessoa
	p.Nome = "Gabriel"
	p.Sobrenome = "Silva"
	p.Idade = 21
	p.Cidade = "São Paulo"

	fmt.Println("Nome: ", p.Nome, "Sobrenome: ", p.Sobrenome, "Idade: ", p.Idade, "Cidade: ", p.Cidade)

	p.Nome = "Zé"
	p.Sobrenome = "Das Couves"
	p.Idade = 30
	p.Cidade = "Recife"

	fmt.Println("Nome: ", p.Nome, "Sobrenome: ", p.Sobrenome, "Idade: ", p.Idade, "Cidade: ", p.Cidade)
}
