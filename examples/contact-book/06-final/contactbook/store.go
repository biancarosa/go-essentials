package contactbook

import (
	"errors"

	"github.com/biancarosa/go-essentials/examples/contact-book/06-final/model"
	"github.com/google/uuid"
)

var (
	ErrContactNotFound      = errors.New("contact not found")
	ErrContactAlreadyExists = errors.New("contact already exists")
)

type Store map[uuid.UUID]model.Contact

func NewStore() Store {
	return make(Store)
}

func Add(store Store, contact model.Contact) error {
	if _, ok := store[contact.ID]; ok {
		return ErrContactAlreadyExists
	}
	store[contact.ID] = contact
	return nil
}

func FindByID(store Store, id uuid.UUID) (model.Contact, error) {
	contact, ok := store[id]
	if !ok {
		return model.Contact{}, ErrContactNotFound
	}
	return contact, nil
}

func FindByEmail(store Store, email string) (model.Contact, error) {
	for _, contact := range store {
		if contact.Email == email {
			return contact, nil
		}
	}
	return model.Contact{}, ErrContactNotFound
}

func List(store Store) []model.Contact {
	contacts := make([]model.Contact, 0, len(store))
	for _, contact := range store {
		contacts = append(contacts, contact)
	}
	return contacts
}

func ListFavorites(store Store) []model.Contact {
	contacts := make([]model.Contact, 0)
	for _, contact := range store {
		if contact.Favorite {
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

func Remove(store Store, id uuid.UUID) error {
	if _, ok := store[id]; !ok {
		return ErrContactNotFound
	}
	delete(store, id)
	return nil
}
