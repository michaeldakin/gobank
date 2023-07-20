package main

import (
	"database/sql"
	"log"

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
	connStr := "postgres://michaeldakin:5wiNuhIVqp8B@ep-nameless-wood-619835.ap-southeast-1.aws.neon.tech/neondb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
