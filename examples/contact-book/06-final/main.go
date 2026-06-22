package main

import (
	"fmt"
	"log"

	"github.com/biancarosa/go-essentials/examples/contact-book/06-final/contactbook"
	"github.com/biancarosa/go-essentials/examples/contact-book/06-final/model"
)

func main() {
	store := contactbook.NewStore()
	contact := model.NewContact("Bia", "bia@example.com", "+55 21 99999-9999", true)

	if err := contactbook.Add(store, contact); err != nil {
		log.Fatal(err)
	}

	found, err := contactbook.FindByID(store, contact.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s <%s> [%s]\n", found.Name, found.Email, found.ID)
	fmt.Println("Favorites:", len(contactbook.ListFavorites(store)))
}
