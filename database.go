package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type DatabaseStore struct {
	db *sql.DB
}

func NewDatabaseStore() (*DatabaseStore, error) {
	dbfile := "gobank.db"
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		log.Fatal("Failed to open database :", err)
	}
	fmt.Printf("Successfully read database file\n")

	return &DatabaseStore{
		db: db,
	}, nil
}

func (s *DatabaseStore) Init() error {
	return s.createAccountsTable()
}

func (s *DatabaseStore) createAccountsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE,
			number REAL UNIQUE,
			balance REAL,
			created_at DATETIME,
			last_updated DATETIME DEFAULT NULL
		)`

	_, err := s.db.Exec(query)
	return err
}

// CreateAccount func - POST to create account in sqlite
func (s *DatabaseStore) CreateAccount(acc *Account) error {
	// POST /account
	// Basic account creation with POST request
	// ID will autoincrement
	// Email must be unique
	// Number must be unique
	query := `
		INSERT INTO accounts (first_name, last_name, email, number, balance, created_at, last_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	resp, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
		acc.UpdatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

// PATCH /account
func (s *DatabaseStore) UpdateAccount(*Account) error {
	return nil
}

// DELETE /account/:id
// This is a HARD delete of data
// TODO: create soft delete functionality - is_active or is_deleted column?
func (s *DatabaseStore) DeleteAccount(id int) error {
	_, err := s.db.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	return err
}

// GET /account
func (s *DatabaseStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`SELECT * FROM accounts`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// GET /accounts/:id
func (s *DatabaseStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query(`SELECT * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("Account %d not found", id)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	return account, err
}
