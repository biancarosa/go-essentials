package main

import (
	"fmt"
	"log"
)

type Contact struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Favorite bool
}

func addContact(contacts []Contact, contact Contact) []Contact {
	return append(contacts, contact)
}

func findContactByID(contacts []Contact, id string) (Contact, error) {
	for _, contact := range contacts {
		if contact.ID == id {
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
		ID:       "550e8400-e29b-41d4-a716-446655440000",
		Name:     "Bia",
		Email:    "bia@example.com",
		Phone:    "+55 21 99999-9999",
		Favorite: true,
	})

	contact, err := findContactByID(contacts, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contact.Name)
	fmt.Println("Favorites:", countFavorites(contacts))
}
