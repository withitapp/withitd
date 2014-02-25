package vldte

import (
	"errors"
)

func Less(i int, limit int) error {
	if i >= limit {
		return errors.New("i is greater than or equal to limit.")
	}
	return nil
}

func Greater(i int, limit int) error {

	if i <= limit {
		return errors.New("i less than or equal to limit.")
	}
	return nil

}
