package main

import (
	"errors"
	"time"
)

type UserTable struct{}

func NewUserTable() *UserTable {
	return &UserTable{}
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
