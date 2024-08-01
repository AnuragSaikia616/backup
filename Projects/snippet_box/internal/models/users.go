package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := "INSERT INTO users (name, email, hashed_password, created) VALUES(?,?,?,UTC_TIMESTAMP())"
	_, err = m.DB.Exec(query, name, email, hash)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) authenticate(email, password string) error {
	return nil
}

func (m *UserModel) exists(id int) error {
	return nil
}
