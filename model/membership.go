package model

import (
	"time"
)

//Struct to represent user memberships in polls
type Membership struct {
	ID        int       "json:`id`"
	CreatedAt time.Time "json:`created_at`"
	UpdatedAt time.Time "json:`updated_at`"
	UserID    int       "json:`user_id`"
	PollID    int       "json:`poll_id`"
	Response  bool      "json:`response`"
}
