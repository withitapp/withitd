package main

import (
	"database/sql"
	"errors"
)

type MembershipTable struct {
	db *sql.DB
}

func NewMembershipTable(db *sql.DB) *MembershipTable {
	return &MembershipTable{db}
}

// BUG Implement
func (t *MembershipTable) Select(id int) (*Membership, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) SelectAll(id []int) (*[]Membership, error) {
	return nil, errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) Insert(membership *Membership) (int, error) {
	return 0, errors.New("not implemented")
}

// BUG Implement
func (t *Membership) Update(membership *Membership) error {
	return errors.New("not implemented")
}

// BUG Implement
func (t *MembershipTable) Remove(membership *Membership) error {
	return errors.New("not implemented")
}
