package main

import (
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

func (u *User) Validate() []error {

}
