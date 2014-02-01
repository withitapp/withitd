package main

import (
	"time"
)

// Struct to represent a friendship between users, Alpha and Beta.
type Friendship struct {
	ID        int
	CreatedAt time.Time
	AlphaID   int
	BetaID    int
}
