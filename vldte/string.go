package vldte

import (
	"errors"
	"regexp"
)

func Email(str string) error {
	var validEmail = regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`)

	if !validEmail.MatchString(str) {
		return errors.New("string is not an email.")
	}
	return nil
}

func Presence(str string) error {
	var validPresence = regexp.MustCompile(`^$`)

	if validPresence.MatchString(str) {
		return errors.New("string is empty.")
	}
	return nil
}

func Length(str string, less int, greater int) error {
	//TODO should this be >= <= instead of > <?
	if len(str) > greater || len(str) < less {
		return errors.New("string is invalid.")
	}
	return nil
}

func Numeric(str string) error {
	var validNumeric = regexp.MustCompile(`^\d+$`)

	if !validNumeric.MatchString(str) {
		return errors.New("string is not numeric.")
	}
	return nil
}

func AlphaNumeric(str string) error {
	var validNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if !validNumeric.MatchString(str) {
		return errors.New("string is not alphanumeric.")
	}
	return nil
}
