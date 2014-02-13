package main

import (
	"database/sql"
	"errors"
	"github.com/withitapp/withitd/model"
)

type MembershipTable struct {
	db *sql.DB
}

func NewMembershipTable(db *sql.DB) *MembershipTable {
	return &MembershipTable{db}
}

// BUG Implement
func (t *MembershipTable) Select(id int) (*model.Membership, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) SelectAll(id []int) (*[]model.Membership, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) Insert(membership *model.Membership) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) Update(membership *model.Membership) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) Remove(membership *model.Membership) error {
	return errors.New("not implemented")
}
