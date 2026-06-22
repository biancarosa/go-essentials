package model

import "github.com/google/uuid"

type Contact struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Phone    string
	Favorite bool
}

func NewContact(name, email, phone string, favorite bool) Contact {
	return Contact{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		Favorite: favorite,
	}
}
