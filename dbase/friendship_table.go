package dbase

import (
	"database/sql"
	"errors"
	"github.com/withitapp/withitd/model"
)

type FriendshipTable struct {
	db *sql.DB
}

func NewFriendshipTable(db *sql.DB) *FriendshipTable {
	return &FriendshipTable{db}
}

// BUG Implement
func (t *FriendshipTable) Select(id int) (*model.Friendship, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) SelectAll(id []int) (*[]model.Friendship, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Insert(friendship *model.Friendship) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Update(friendship *model.Friendship) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *FriendshipTable) Remove(friendship *model.Friendship) error {
	return errors.New("not implemented")
}
