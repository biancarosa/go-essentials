package main

import "fmt"

func main() {
	const appName = "Contact Book"

	contactID := "550e8400-e29b-41d4-a716-446655440000"
	contactName := "Bia"
	contactEmail := "bia@example.com"
	contactPhone := "+55 21 99999-9999"
	contactFavorite := true

	fmt.Println(appName)
	fmt.Println("ID:", contactID)
	fmt.Println("Name:", contactName)
	fmt.Println("Email:", contactEmail)
	fmt.Println("Phone:", contactPhone)
	fmt.Println("Favorite:", contactFavorite)
}
