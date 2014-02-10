package main

import (
	"database/sql"
	"errors"
)

type FriendshipTable struct {
	db *sql.DB
}

func NewFriendshipTable(db *sql.DB) *FriendshipTable {
	return &FriendshipTable{db}
}

// BUG Implement
func (t *FriendshipTable) Select(id int) (*Friendship, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) SelectAll(id []int) (*[]Friendship, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Insert(friendship *Friendship) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Update(friendship *Friendship) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Remove(friendship *Friendship) error {
	return errors.New("not implemented")
}
