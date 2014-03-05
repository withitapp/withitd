package model

import (
	"time"
)

// Struct to represent a friendship between users, Alpha and Beta.
type Friendship struct {
	ID        int       "json:`id`"
	CreatedAt time.Time "json:`created_at`"
	AlphaID   int       "json:`alpha_id`"
	BetaID    int       "json:`beta_id`"
}
