package main

import "fmt"

const MAX int = 100

func main() {
	// expr := n < MAX // declarando expr e atribuindo o valor (n < MAX)
	// if expr {
	// 	fmt.Printf("%d é menor do que MAX == %d\n", n, MAX)
	// } else {
	// 	fmt.Printf("%d é maior do que MAX == %d\n", n, MAX)
	// }

	for n := 0; n < MAX; n = n + 1 {
		if n%2 == 0 {
			fmt.Printf("%d é par\n", n)
		} else {
			fmt.Printf("%d é impar\n", n)
		}

	}

	// for (inicializacao); (condicao que precisa ser verdade); (incremento)
	for i := 0; i < MAX; i++ {
	}
}
