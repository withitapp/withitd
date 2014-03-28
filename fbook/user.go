package fbook

/*
import (
	"time"
)
*/

type User struct {
	ID  string `json:"id", facebook:"id,required"`
	Bio string `json:"bio", facebook:"bio,required"`
	//Birthday           time.Time `json:"birthday"`
	Birthday           string `json:"birthday", facebook:"birthday,required"`
	Email              string `json:"email", facebook:"email,required"`
	FirstName          string `json:"first_name", facebook: "first_name,required"`
	MiddleName         string `json:"middle_name", facebook:"middle_name,required"`
	LastName           string `json:"last_name", facebook:"last_name,required"`
	Name               string `json:"name", facebook:"name,required"`
	NameFormat         string `json:"name_format", facebook:"name_format,required"`
	Gender             string `json:"gender", facebook:"gender,required"`
	Installed          bool   `json:"installed, facebook:"installed,required"`
	IsVerified         bool   `json:"is_verified", facebook:"is_verified,required"`
	Link               string `json:"link", facebook:"link,required"`
	Locale             string `json:"locale", facebook:"locale,required"`
	Political          string `json:"political", facebook:"political,required"`
	Quotes             string `json:"quotes", facebook:"quotes,required"`
	RelationshipStatus string `json:"relationship_status", facebook:"relationship_status,required"`
	Religion           string `json:"religion", facebook:"religion,required"`
	ThirdPartyID       string `json:"third_party_id", facebook:"third_party_id,required"`
	Username           string `json:"username", facebook:"username,required"`
	Verified           bool   `json:"verified", facebook:"verified,required"`
	Website            string `json:"website", facebook:"website,required"`
}
