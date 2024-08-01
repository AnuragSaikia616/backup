package main

import "time"

type Account struct {
	ID        int       `json: "id"`
	FirstName string    `json: "firstname"`
	LastName  string    `json: "lastname"`
	Balance   float64   `json: "balance"`
	CreatedAt time.Time `json: "createdAt"`
}

func newAccount(firstname string, lastname string) *Account {
	return &Account{
		FirstName: firstname,
		LastName:  lastname,
		Balance:   float64(0),
		CreatedAt: time.Now().UTC(),
	}
}
