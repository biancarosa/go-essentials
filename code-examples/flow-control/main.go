package main

import "fmt"

var x int

func anotherFunction() {
	fmt.Println(x)
	// fmt.Println(y) - não pode ser usado aqui :)
}
func main() {
	var y int
	x = 10
	y = 10
	fmt.Println("Qual o valor de x?", x) // 10
	if x == y {
		for x := 100; ; {
			fmt.Println("Qual o valor de x?", x) // 100
			if x == 100 {
				break
			}
		}
	}
}
