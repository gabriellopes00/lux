package validators

import "helpy/pkg/entities"

type UserValidator interface {
	Validate(user *entities.User) error
}
