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

	if len(values.Get("title")) == 0 {
		return nil, errors.New("title field is empty.")
	}

	if len(values.Get("description")) == 0 {
		return nil, errors.New("description field is empty.")
	}

	if len(values.Get("user_id")) == 0 {
		return nil, errors.New("user_id field is empty.")
	}

	if len(values.Get("ends_at")) == 0 {
		return nil, errors.New("ends_at field is empty.")
	}

	//Convert user ID to string
	userID, err := strconv.Atoi(values.Get("user_id"))
	if err != nil {
		return nil, err
	}

	endsAt, err := time.Parse(time.RFC3339, values.Get("ends_at"))
	if err != nil {
		return nil, err
	}

	return &Poll{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       values.Get("title"),
		Description: values.Get("description"),
		UserID:      userID,
		EndsAt:      endsAt,
	}, nil
}

func (p *Poll) UpdateFromValues(values url.Values) error {

	p.UpdatedAt = time.Now()

	if values.Get("title") == "" {
		return errors.New("title field is empty.")
	}
	p.Title = values.Get("title")

	if len(values.Get("description")) == 0 {
		return errors.New("description field is empty.")
	}
	p.Description = values.Get("description")

	if len(values.Get("user_id")) == 0 {
		return errors.New("user_id field is empty.")
	}

	if len(values.Get("ends_at")) == 0 {
		return errors.New("ends_at field is empty.")
	}

	//Convert user ID to string
	userID, err := strconv.Atoi(values.Get("user_id"))
	if err != nil {
		return err
	}
	p.UserID = userID

	endsAt, err := time.Parse(time.RFC3339, values.Get("ends_at"))
	if err != nil {
		return err
	}
	p.EndsAt = endsAt

	return nil
}

func (p *Poll) Validate() error {
	stringLengthMax := 25
	stringLengthMin := 4

	//Check Title
	return vldte.Length(p.Title, stringLengthMin, stringLengthMax)
}
