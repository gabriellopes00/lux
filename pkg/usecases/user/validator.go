package user

import "helpy/pkg/entities"

type Validator interface {
	Validate(user *entities.User) error
}
