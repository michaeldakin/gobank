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
	fmt.Println("Imported .env file")

	connStr := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
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

func (s *DatabaseStore) CreateAccount(*Account) error {
	return nil
}
