package operations

import (
	"errors"
	"fmt"
)

// Retorna 0 na divisão
// func Divide(a int, b int) int {
// 	fmt.Printf("Dividindo %d / %d\n", a, b)
// 	if b == 0 {
// 		fmt.Println("Não posso fazer divisão por zero!")
// 		return 0
// 	}
// 	return a / b
// }

// panic e para a execução completamente
// func Divide(a int, b int) int {
// 	fmt.Printf("Dividindo %d / %d\n", a, b)
// 	if b == 0 {
// 		panic("Divisão por zero não permitida")
// 	}
// 	return a / b
// }

func Divide(a int, b int) (int, error) {
	fmt.Printf("Dividindo %d / %d\n", a, b)
	if b == 0 {
		err := errors.New("divisão por zero não permitida")
		return 0, err
	}
	return a / b, nil
}
