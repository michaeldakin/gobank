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

	rows, err := db.Query("SELECT * from playing_with_neon")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
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
				number SERIAL,
				balance DECIMAL,
				created_at TIMESTAMP,
				updated_at TIMESTAMP
			)`

	_, err := s.db.Exec(query)
	return err
}

func (s *DatabaseStore) CreateAccount(*Account) error {
	return nil
}
func (s *DatabaseStore) UpdateAccount(*Account) error {
	return nil
}
func (s *DatabaseStore) DeleteAccount(id int) error {
	return nil
}
func (s *DatabaseStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
