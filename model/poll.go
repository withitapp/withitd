package model

import (
	"time"
)

// Struct to represent a poll for the app
type Poll struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	UserID      int
	EndsAt      time.Time
}

//TODO Implement
func (p *Poll) Validate() []error {
	return nil
}
