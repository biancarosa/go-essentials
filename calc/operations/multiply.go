package operations

import "fmt"

// funções que começam com letra minuscula == privada
// só são visiveis no mesmo pacote
func multiply(a int, b int) int {
	fmt.Printf("Multiplicando %d * %d\n", a, b)
	return a * b
}

// Multiply com M maisculo é uma função "exportada"
func Multiply(a int, b int) int {
	return multiply(a, b)
}
