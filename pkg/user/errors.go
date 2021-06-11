package user

import "errors"

var (
	ErrExistingEmail = errors.New("email already in use")
)
