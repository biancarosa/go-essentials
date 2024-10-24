package main

import (
	"calculadora/operations"
	"fmt"
	"os"
	"strconv"
)

// Coisas futuras legais de fazer:
// 1 - Passar a tratar valores de ponto flutuante: go run . "*" 2.5 5
// 2 - Multiplos (infinitos ou variáveis valores de entrada):
// https://www.digitalocean.com/community/tutorials/how-to-use-variadic-functions-in-go
// ---- mesmo programa sem modificação funcionaria para
// ----- go run . "+" 1000 500 100 1000
// ----- go run . "+" 1000 500 100 1000 100
// ----- go run . "+" 1000 500 100 1000 100 20
// 3 - Receber uma entrada mais bonitinha e tratar ela: go run . =100/10
func main() {

	// slice (array? maizomenos) uma lista de valores.
	// args := os.Args
	// fmt.Println(args)

	op := os.Args[1]

	s1 := os.Args[2]
	s2 := os.Args[3]

	n1, err := strconv.Atoi(s1)
	if err != nil {
		panic("Valor inválido")
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		panic("Valor inválido")
	}

	fmt.Printf("Executando a operação %s com os valores %d e %d\n", op, n1, n2)
	var result int
	switch op {
	case "+":
		result = operations.Sum(n1, n2)
	case "*":
		result = operations.Multiply(n1, n2)
	case "-":
		result = operations.Subtract(n1, n2)
	case "/":
		result, err = operations.Divide(n1, n2)
	}
	// verificar se o erro existe antes de tentar usar o valor de result
	if err != nil {
		panic(err)
		// panic para a execução do programa
	}

	fmt.Println("Resultado === ", result)

}
