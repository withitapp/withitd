package dbase

import (
	"database/sql"
	"errors"
	"github.com/withitapp/withitd/model"
)

type PollTable struct {
	db *sql.DB
}

func NewPollTable(db *sql.DB) *PollTable {
	return &PollTable{db}
}

// BUG Implement
func (t *PollTable) Select(id int) (*model.Poll, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) SelectAll(id []int) (*[]model.Poll, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Insert(poll *model.Poll) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Update(poll *model.Poll) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *PollTable) Remove(poll *model.Poll) error {
	return errors.New("not implemented")
}
