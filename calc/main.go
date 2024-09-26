package main

import "fmt"

func sum(a int, b int) int {
	fmt.Printf("Somando %d + %d\n", a, b)
	return a + b
}

func multiply(a int, b int) int {
	fmt.Printf("Multiplicando %d * %d\n", a, b)
	return a * b
}

func divide(a int, b int) int {
	fmt.Printf("Dividindo %d / %d\n", a, b)
	return a / b
}

func subtract(a int, b int) int {
	fmt.Printf("Subtraindo %d - %d\n", a, b)
	return a - b
}

func main() {
	fmt.Println("Calculadora #GoEssentials")

	result := sum(5, 10) // criação e atribuição
	fmt.Println(result)

	result = divide(10, 5) // só atribuindo
	fmt.Println(result)

	result = multiply(5, 100) // só atribuindo
	fmt.Println(result)

	result = subtract(50, 10) // só atribuindo
	fmt.Println(result)

}
