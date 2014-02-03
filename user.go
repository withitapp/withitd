package main

import (
	"errors"
	"time"
)

// Struct to represent the user of the app and their associated data
type User struct {
	ID         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Username   string
	Email      string
	FirstName  string
	LastName   string
	FacebookID string
}

// Selects a user row using the given ID field
// BUG Returns a dummy user
func SelectUser(id int) (*User, error) {
	if id == 0 {
		return nil, errors.New("Not valid User.ID")
	}

	return &User{
		ID:         id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Username:   "jdoe",
		Email:      "john@doe.me",
		FirstName:  "John",
		LastName:   "Doe",
		FacebookID: "13450134143249",
	}, nil
}

// TODO Implement SelectUsers(ids []int) ([]*User, error)

// TODO Implement InsertUser(user *User) (int, error)

// TODO Implement UpdateUser(user *User) error

// TODO Implement RemoveUser(user *User) error

// TODO Implement User.FriendshipIDs() []int

// TODO Implement User.InvitationIDs() []int

// TODO Implement User.MembershipIDs() []int
