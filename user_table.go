package main

import (
	"database/sql"
	"errors"
	"github.com/withitapp/withitd/model"
)

type UserTable struct {
	db *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
	return &UserTable{db}
}

// BUG Implement
func (t *UserTable) Select(id int) (*model.User, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) SelectAll(id []int) (*[]model.User, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Insert(user *model.User) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Update(user *model.User) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Remove(user *model.User) error {
	return errors.New("not implemented")
}
