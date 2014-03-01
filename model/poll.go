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

func (p *Poll) Validate() error {
	stringLengthMax := 25
	stringLengthMin := 4

	//Check Title
	return vldte.Length(u.Username, stringLengthMin, stringLengthMax)
}
