package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	fmt.Println("Imported database .env file")

	connStr := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connected to database: %s\n", connStr)
	}

	//rows, err := db.Query("SELECT * from playing_with_neon")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(rows)

	return &DatabaseStore{
		db: db,
	}, nil
}

func (s *DatabaseStore) Init() error {
	return s.createAccountTable()
}

func (s *DatabaseStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
				id SERIAL PRIMARY KEY,
				first_name TEXT,
				last_name TEXT,
				email TEXT UNIQUE,
				number SERIAL UNIQUE,
				balance DECIMAL,
				created_at TIMESTAMP,
				last_updated TIMESTAMP DEFAULT NULL
			)`

	_, err := s.db.Exec(query)
	return err
}

func (s *DatabaseStore) CreateAccount(acc *Account) error {
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

func (s *DatabaseStore) UpdateAccount(*Account) error {
	return nil
}

func (s *DatabaseStore) DeleteAccount(id int) error {
	return nil
}

func (s *DatabaseStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
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
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *DatabaseStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
