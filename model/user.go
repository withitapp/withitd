package model

import (
	"github.com/withitapp/withitd/vldte"
	"time"
)

// Struct to represent the user of the app and their associated data
type User struct {
	ID         int       "json:`id`"
	CreatedAt  time.Time "json:`created_at`"
	UpdatedAt  time.Time "json:`updated_at`"
	Username   string    "json:`username`"
	Email      string    "json:`email`"
	FirstName  string    "json:`first_name`"
	LastName   string    "json:`last_name`"
	FacebookID string    "json:`facebook_id`"
}

func (u *User) Validate() error {
	stringLengthMax := 25
	stringLengthMin := 1

	//Check Username
	err := vldte.AlphaNumeric(u.Username)
	if err != nil {
		return err
	}

	err = vldte.Length(u.Username, stringLengthMin, stringLengthMax)
	if err != nil {
		return err
	}

	//Check Email
	err = vldte.Email(u.Email)
	if err != nil {
		return err
	}

	//Check First Name
	err = vldte.Alphabetic(u.FirstName)
	if err != nil {
		return err
	}

	err = vldte.Length(u.FirstName, stringLengthMin, stringLengthMax)
	if err != nil {
		return err
	}

	//Check Last Name
	err = vldte.Alphabetic(u.LastName)
	if err != nil {
		return err
	}

	err = vldte.Length(u.LastName, stringLengthMin, stringLengthMax)
	if err != nil {
		return err
	}

	//Check Facebook ID
	err = vldte.Numeric(u.FacebookID)
	if err != nil {
		return err
	}

	return nil
}
