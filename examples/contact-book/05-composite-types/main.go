package main

import (
	"errors"
	"fmt"
	"log"
)

var errContactNotFound = errors.New("contact not found")

type Contact struct {
	Name     string
	Email    string
	Phone    string
	Favorite bool
}

type Store map[string]Contact

func addContact(store Store, contact Contact) {
	store[contact.Email] = contact
}

func findContactByEmail(store Store, email string) (Contact, error) {
	contact, ok := store[email]
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

func removeContactByEmail(store Store, email string) error {
	if _, ok := store[email]; !ok {
		return errContactNotFound
	}
	delete(store, email)
	return nil
}

func main() {
	store := make(Store)
	contact := Contact{
		Name:     "Bia",
		Email:    "bia@example.com",
		Phone:    "+55 21 99999-9999",
		Favorite: true,
	}
	addContact(store, contact)

	found, err := findContactByEmail(store, contact.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(found.Name)
	fmt.Println("Total:", len(listContacts(store)))
}
