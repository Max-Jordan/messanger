// Package storage
package storage

import (
	"database/sql"

	"github.com/Max-Jordan/messanger/internal/models"
)

const initPGstorage = `
create table if not exists (
id int )`

type UserRepository interface {
	Save(models.User) error
	GetByID(id int) (models.User, error)
}

type userPostgreStorage struct {
	db *sql.DB
}

func NewUserPostgreStorage(db *sql.DB) *userPostgreStorage {
	return &userPostgreStorage{db: db}
}

func (s *userPostgreStorage) Save(u models.User) error {
	return nil
}

func (s *userPostgreStorage) GetByID(id int) (models.User, error) {
	return models.User{}, nil
}


