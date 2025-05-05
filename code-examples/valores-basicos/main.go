// Exemplo de valores básicos em Go
package main

import (
	"fmt"
)

func main() {
	// Strings
	fmt.Println("== Strings ==")
	fmt.Println("hello" + "world")

	// Concatenação com espaço
	fmt.Println("hello" + " " + "world")

	// Inteiros
	fmt.Println("\n== Inteiros ==")
	fmt.Println("2 + 6 =", 2+6)
	fmt.Println("10 - 2 =", 10-2)
	fmt.Println("4 * 3 =", 4*3)
	fmt.Println("8 / 4 =", 8/4)

	// Ponto flutuante
	fmt.Println("\n== Ponto Flutuante ==")
	fmt.Println("10.0 / 3.0 =", 10.0/3.0)
	fmt.Println("1.5 * 2.0 =", 1.5*2.0)

	// Booleanos
	fmt.Println("\n== Booleanos ==")
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// Comparações
	fmt.Println("\n== Comparações ==")
	fmt.Println("5 > 3 =", 5 > 3)
	fmt.Println("5 < 3 =", 5 < 3)
	fmt.Println("5 == 5 =", 5 == 5)
	fmt.Println("5 != 5 =", 5 != 5)
}
