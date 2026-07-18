package main

import "fmt"

func main() {
	const appName = "Contact Book"

	contactName := "Bia"
	contactEmail := "bia@example.com"
	contactPhone := "+55 21 99999-9999"
	contactFavorite := true

	fmt.Println(appName)
	fmt.Println("Name:", contactName)
	fmt.Println("Email:", contactEmail)
	fmt.Println("Phone:", contactPhone)
	fmt.Println("Favorite:", contactFavorite)
}
