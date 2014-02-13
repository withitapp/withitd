package model

import (
	"time"
)

//Struct to represent user memberships in polls
type Membership struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	PollID    int
	Response  bool
}
