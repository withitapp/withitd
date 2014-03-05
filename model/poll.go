package model

import (
	"time"
)

// Struct to represent a poll for the app
type Poll struct {
	ID          int       "json:`id`"
	CreatedAt   time.Time "json:`created_at`"
	UpdatedAt   time.Time "json:`updated_at`"
	Title       string    "json:`title`"
	Description string    "json:`description`"
	UserID      int       "json:`user_id`"
	EndsAt      time.Time "json:`ends_at`"
}

func (p *Poll) Validate() error {
	stringLengthMax := 25
	stringLengthMin := 4

	//Check Title
	return vldte.Length(u.Username, stringLengthMin, stringLengthMax)
}
