package fbook

type User struct {
	ID                 string `json:"id" facebook:"id,required"`
	Bio                string `json:"bio" facebook:"bio"`
	Birthday           string `json:"birthday" facebook:"birthday"`
	Email              string `json:"email" facebook:"email,required"`
	FirstName          string `json:"first_name" facebook: "first_name,required"`
	MiddleName         string `json:"middle_name" facebook:"middle_name"`
	LastName           string `json:"last_name" facebook:"last_name,required"`
	Name               string `json:"name" facebook:"name"`
	NameFormat         string `json:"name_format" facebook:"name_format"`
	Gender             string `json:"gender" facebook:"gender"`
	Installed          bool   `json:"installed" facebook:"installed"`
	IsVerified         bool   `json:"is_verified" facebook:"is_verified"`
	Link               string `json:"link" facebook:"link"`
	Locale             string `json:"locale" facebook:"locale"`
	Political          string `json:"political" facebook:"political"`
	Quotes             string `json:"quotes" facebook:"quotes"`
	RelationshipStatus string `json:"relationship_status" facebook:"relationship_status"`
	Religion           string `json:"religion" facebook:"religion"`
	ThirdPartyID       string `json:"third_party_id" facebook:"third_party_id"`
	Username           string `json:"username" facebook:"username,required"`
	Verified           bool   `json:"verified" facebook:"verified"`
	Website            string `json:"website" facebook:"website"`
}
