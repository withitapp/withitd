package main

import (
	"database/sql"
	"errors"
)

type UserTable struct {
	db *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
	return &UserTable{db}
}

// BUG Implement
func (t *UserTable) Select(id int) (*User, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) SelectAll(id []int) (*[]User, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Insert(user *User) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Update(user *User) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *UserTable) Remove(user *User) error {
	return errors.New("not implemented")
}
