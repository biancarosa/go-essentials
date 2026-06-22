package main

import (
	"errors"
	"fmt"
	"log"
)

var errContactNotFound = errors.New("contact not found")

type Contact struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Favorite bool
}

type Store map[string]Contact

func addContact(store Store, contact Contact) {
	store[contact.ID] = contact
}

func findContactByID(store Store, id string) (Contact, error) {
	contact, ok := store[id]
	if !ok {
		return Contact{}, errContactNotFound
	}
	return contact, nil
}

func listContacts(store Store) []Contact {
	contacts := make([]Contact, 0, len(store))
	for _, contact := range store {
		contacts = append(contacts, contact)
	}
	return contacts
}

func removeContactByID(store Store, id string) error {
	if _, ok := store[id]; !ok {
		return errContactNotFound
	}
	delete(store, id)
	return nil
}

func main() {
	store := make(Store)
	contact := Contact{
		ID:       "550e8400-e29b-41d4-a716-446655440000",
		Name:     "Bia",
		Email:    "bia@example.com",
		Phone:    "+55 21 99999-9999",
		Favorite: true,
	}
	addContact(store, contact)

	found, err := findContactByID(store, contact.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(found.Name)
	fmt.Println("Total:", len(listContacts(store)))
}
