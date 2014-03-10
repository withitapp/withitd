package fbook

/*
import (
	"time"
)
*/

type User struct {
	ID  string `json:"id"`
	Bio string `json:"bio"`
	//Birthday           time.Time `json:"birthday"`
	Birthday           string `json:"birthday"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	MiddleName         string `json:"middle_name"`
	LastName           string `json:"last_name"`
	Name               string `json:"name"`
	NameFormat         string `json:"name_format"`
	Gender             string `json:"gender"`
	Installed          bool   `json:"installed"`
	IsVerified         bool   `json:"is_verified"`
	Link               string `json:"link"`
	Locale             string `json:"locale"`
	Political          string `json:"political"`
	Quotes             string `json:"quotes"`
	RelationshipStatus string `json:"relationship_status"`
	Religion           string `json:"religion"`
	ThirdPartyID       string `json:"third_party_id"`
	Username           string `json:"username"`
	Verified           bool   `json:"verified"`
	Website            string `json:"website"`
}
