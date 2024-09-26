package main

import (
	"fmt"

	"calculadora/operations"
)

func main() {
	fmt.Println("Calculadora #GoEssentials")

	result := operations.Sum(5, 10) // criação e atribuição
	fmt.Println(result)

	result = operations.Divide(10, 5) // só atribuindo
	fmt.Println(result)

	result = operations.Multiply(5, 100) // só atribuindo
	fmt.Println(result)

	result = operations.Subtract(50, 10) // só atribuindo
	fmt.Println(result)

}
