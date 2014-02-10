package main

import (
	"database/sql"
	"errors"
)

type PollTable struct {
	db *sql.DB
}

func NewPollTable(db *sql.DB) *PollTable {
	return &PollTable{db}
}

// BUG Implement
func (t *PollTable) Select(id int) (*Poll, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) SelectAll(id []int) (*[]Poll, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Insert(poll *Poll) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Update(poll *Poll) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Remove(poll *Poll) error {
	return errors.New("not implemented")
}
