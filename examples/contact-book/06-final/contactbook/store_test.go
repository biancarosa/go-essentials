package contactbook

import (
	"errors"
	"testing"

	"github.com/biancarosa/go-essentials/examples/contact-book/06-final/model"
	"github.com/google/uuid"
)

func TestAddAndFindByID(t *testing.T) {
	store := NewStore()
	contact := model.NewContact("Bia", "bia@example.com", "+55 21 99999-9999", true)

	if err := Add(store, contact); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	found, err := FindByID(store, contact.ID)
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if found != contact {
		t.Fatalf("FindByID() = %#v, want %#v", found, contact)
	}
}

func TestAddRejectsDuplicateID(t *testing.T) {
	store := NewStore()
	contact := model.NewContact("Bia", "bia@example.com", "+55 21 99999-9999", true)

	if err := Add(store, contact); err != nil {
		t.Fatalf("first Add() error = %v", err)
	}
	if err := Add(store, contact); !errors.Is(err, ErrContactAlreadyExists) {
		t.Fatalf("second Add() error = %v, want %v", err, ErrContactAlreadyExists)
	}
}

func TestFindByIDReturnsNotFound(t *testing.T) {
	_, err := FindByID(NewStore(), uuid.New())
	if !errors.Is(err, ErrContactNotFound) {
		t.Fatalf("FindByID() error = %v, want %v", err, ErrContactNotFound)
	}
}
