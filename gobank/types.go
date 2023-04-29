package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	uuid := uuid.New()
	return &Account{
		ID:        int(uuid.ID()),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
	}
}
