package main

import "fmt"

type Contact struct {
	Name     string
	Phone    string
	Favorite bool
}

func main() {
	contacts := []Contact{
		{Name: "Bia", Phone: "+55 21 99999-9999", Favorite: true},
		{Name: "Ana", Phone: "", Favorite: false},
		{Name: "Joao", Phone: "+55 11 99999-9999", Favorite: true},
	}

	emptyPhones := 0
	for _, contact := range contacts {
		if contact.Favorite {
			fmt.Println("Favorite:", contact.Name)
		}
		if contact.Phone == "" {
			emptyPhones++
		}
	}

	fmt.Println("Contacts without phone:", emptyPhones)

	switch len(contacts) {
	case 0:
		fmt.Println("empty")
	case 1, 2, 3, 4, 5:
		fmt.Println("starting")
	default:
		fmt.Println("growing")
	}
}
