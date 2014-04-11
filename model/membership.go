package model

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

//Struct to represent user memberships in polls
type Membership struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
	PollID    int       `json:"poll_id"`
	Response  bool      `json:"response"`
}

func NewMembershipFromValues(values url.Values) (*Membership, error) {

	if len(values.Get("user_id")) == 0 {
		return nil, errors.New("user_id field is empty")
	}

	if len(values.Get("poll_id")) == 0 {
		return nil, errors.New("poll_id field is empty")
	}

	if len(values.Get("response")) == 0 {
		return nil, errors.New("response field is empty")
	}

	//Convert ID numbers to string
	userID, err := strconv.Atoi(values.Get("user_id"))
	if err != nil {
		return nil, err
	}

	pollID, err := strconv.Atoi(values.Get("poll_id"))
	if err != nil {
		return nil, err
	}

	pollResponse := false
	if values.Get("response") == "true" {
		pollResponse = true
	}

	return &Membership{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		PollID:    pollID,
		Response:  pollResponse,
	}, nil
}
