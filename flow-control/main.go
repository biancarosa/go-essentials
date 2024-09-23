package main

import "fmt"

const MAX int = 100

func main() {

	n := 0
	expr := n < MAX // declarando expr e atribuindo o valor (n < MAX)
	if expr {
		fmt.Printf("%d é menor do que MAX == %d\n", n, MAX)
	} else {
		fmt.Printf("%d é maior do que MAX == %d\n", n, MAX)
	}

	for n := 0; n < MAX; n++ {
		if n%2 == 0 {
			fmt.Printf("%d é par\n", n)
		} else {
			fmt.Printf("%d é impar\n", n)
		}
	}

	// for (inicializacao); (condicao que precisa ser verdade); (incremento)
	for i := 0; i < MAX; i++ {}

	for {
		n = n+1;
		fmt.Println(n)
		if n > MAX {
			break;
		}
	} // loop infinito (não existe while em Go, apenas for infinito)

	switch n < MAX {
	case true: {
		fmt.Println("n é menor do que max")	
	}	
	case false: {
		fmt.Println("n não é menor do que max")	
	}	
	}

	switch n {
	case 100: fmt.Println("hello")
	default: fmt.Println("não é 100")
	}

}
