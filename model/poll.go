package model

import (
	"errors"
	"github.com/withitapp/withitd/vldte"
	"net/url"
	"strconv"
	"time"
)

// Struct to represent a poll for the app
type Poll struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	EndsAt      time.Time `json:"ends_at"`
}

//Create a Poll object from URL values
func NewPollFromValues(values url.Values) (*Poll, error) {

	if len(values["title"]) == 0 {
		return nil, errors.New("title field is empty.")
	}

	if len(values["description"]) == 0 {
		return nil, errors.New("description field is empty.")
	}

	if len(values["user_id"]) == 0 {
		return nil, errors.New("user_id is empty.")
	}

	//Convert user ID to string
	userID, err := strconv.Atoi(values["user_id"][0])
	if err != nil {
		return nil, err
	}

	return &Poll{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       values["title"][0],
		Description: values["description"][0],
		UserID:      userID,
		//EndsAt: value["ends_at"][0],
	}, nil
}

func (p *Poll) Validate() error {
	stringLengthMax := 25
	stringLengthMin := 4

	//Check Title
	return vldte.Length(p.Title, stringLengthMin, stringLengthMax)
}
