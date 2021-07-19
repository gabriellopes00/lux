package validators

import "lux/pkg/entities"

type UserValidator interface {
	Validate(user *entities.User) error
}
