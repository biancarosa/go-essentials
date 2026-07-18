package main

import (
	"fmt"
	"log"
)

type Contact struct {
	Name     string
	Email    string
	Phone    string
	Favorite bool
}

func addContact(contacts []Contact, contact Contact) []Contact {
	return append(contacts, contact)
}

func findContactByEmail(contacts []Contact, email string) (Contact, error) {
	for _, contact := range contacts {
		if contact.Email == email {
			return contact, nil
		}
	}
	return Contact{}, fmt.Errorf("contact not found")
}

func countFavorites(contacts []Contact) int {
	count := 0
	for _, contact := range contacts {
		if contact.Favorite {
			count++
		}
	}
	return count
}

func main() {
	contacts := []Contact{}
	contacts = addContact(contacts, Contact{
		Name:     "Bia",
		Email:    "bia@example.com",
		Phone:    "+55 21 99999-9999",
		Favorite: true,
	})

	contact, err := findContactByEmail(contacts, "bia@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contact.Name)
	fmt.Println("Favorites:", countFavorites(contacts))
}
