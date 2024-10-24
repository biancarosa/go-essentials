package main

// programa escopo 1 (inicial / maior escopo)

import "fmt"

const i float64 = 5.5
const timeToSleepInSeconds = 600

// main - escopo 2 (menor / mais limitado)
func main() {
	var z, a, b int = 1, 2, 3 // criação/declaração de n variáveis
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	c, d := 3, 5
	fmt.Println(c)
	fmt.Println(d)

	var e, f int
	e = 1
	fmt.Println(e)
	fmt.Println(f)

	// estaticamente tipado + fortemente tipado
	e = 5

	var g bool
	g = true
	fmt.Println(g)

	var h string
	h = "5"
	fmt.Println(h)

	h = "sou uma string h"
	fmt.Println(h)

	fmt.Println(i)

	const pi = 3.14159
	fmt.Println(pi)

	var y string
	y = "hello world alterado"
	helloWorld(y)

	fmt.Println(timeToSleepInSeconds)
}

func helloWorld(y string) {
	x := "x: Hello world!" // criação/declaração e atribuição de x -> :=
	fmt.Println(x)         // referencia a x

	y = "y: Hello World"
	fmt.Println(y) // referencia a y

	const pi = 3.14159
	fmt.Println(pi)
	fmt.Println(timeToSleepInSeconds)
}
