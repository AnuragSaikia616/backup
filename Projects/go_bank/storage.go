package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccountByID(int) (*Account, error)
	CreateAccount(string, string) (int, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) error
}

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore(psqlInfo string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()

}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS ACCOUNTS(
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  balance FLOAT,
  created_at TIMESTAMP
  )`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := "SELECT * FROM ACCOUNTS WHERE id = $1"
	row := s.DB.QueryRow(query, id)
	account := &Account{}
	err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Balance, &account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil

}
func (s *PostgresStore) CreateAccount(firstname, lastname string) (int, error) {

	query := `INSERT INTO ACCOUNTS (
  first_name,
  last_name,
  balance,
  created_at
  ) VALUES (
  $1, $2, $3, $4
  )`
	_, err := s.DB.Exec(query, firstname, lastname, 0, time.Now())
	if err != nil {
		return 0, err
	}
	account := newAccount(firstname, lastname)
	return account.ID, nil

}
func (s *PostgresStore) DeleteAccount(id int) error {
	query := `DELETE FROM ACCOUNTS WHERE id = $1`
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil

}
func (s *PostgresStore) UpdateAccount(account *Account) error {
	query := `UPDATE ACCOUNTS SET first_name = $1, last_name = $2, balance = $3, created_at = $4 WHERE id = $5`
	_, err := s.DB.Exec(query, account.FirstName, account.LastName, account.Balance, account.CreatedAt, account.ID)
	if err != nil {
		return err
	}
	return nil
}
