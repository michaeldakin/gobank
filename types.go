package main

import (
	"math/rand"
	"time"
)

// TransferRequest struct
// Function requires WHO the AMOUNT is going to
type TransferRequest struct {
	AccountNumber int     `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}

// CreateAccountRequest struct
// Basic requirements to make an account
// firstName
// lastName
// email
type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Account struct
// Required fields to create an Account
// TODO add IsActive field for soft deletion
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"last_updated"`
}

// NewAccount function
// Creates account with provided CreateAccountRequest struct
// Randomly generates account Number, inserts timestamp into
// CreatedAt and UpdatedAt
// UpdatedAt will be updated anytime an account is modified
func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Number:    int64(rand.Intn(100000)), //account number
		IsActive:  1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
