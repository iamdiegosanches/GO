package main

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publicationDate"`
	Publisher       string    `json:"publisher"`
}

func NewBook(name string, author string, publicationDate time.Time, publisher string) *Book {
	return &Book{
		ID:              uuid.New(),
		Name:            name,
		Author:          author,
		PublicationDate: publicationDate,
		Publisher:       publisher,
	}
}
