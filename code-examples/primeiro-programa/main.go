// Primeiro programa em Go
package main

import (
	"fmt"
)

func main() {
	// Exibe uma mensagem de boas-vindas
	fmt.Println("Olá, mundo! Este é meu primeiro programa em Go!")

	// Exemplos de variáveis
	nome := "Estudante"
	idade := 25

	// Usando formatação para exibir valores
	fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)

	// Calculando e exibindo resultados
	a, b := 10, 5
	soma := a + b
	fmt.Println("A soma de", a, "e", b, "é", soma)
}
