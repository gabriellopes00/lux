package validation

import (
	"fmt"
	"lux/pkg/entities"

	"github.com/asaskevich/govalidator"
)

type UserGoValidator struct{}

func (v UserGoValidator) Validate(data *entities.User) error {
	if !govalidator.IsEmail(data.Email) {
		return fmt.Errorf("invalid field %s", "email")

	} else if !govalidator.IsURL(data.AvatarUrl) {
		return fmt.Errorf("invalid field %s", "avatar_url")

	}

	return nil
}
