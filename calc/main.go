package main

import "fmt"

func sum(a int, b int) int {
	fmt.Printf("Somando %d + %d\n", a, b)
	return a + b
}

// func multiply() {}
// func divide()   {}
// func subtract() {}

func main() {
	fmt.Println("Calculadora #GoEssentials")

	result := sum(5, 10)
	fmt.Println(result)

}
