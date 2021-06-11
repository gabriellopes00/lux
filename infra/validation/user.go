package validation

import (
	"fmt"
	"helpy/pkg/entities"

	"github.com/asaskevich/govalidator"
)

func ErrInvalidField(field string) error {
	return fmt.Errorf("invalid field %s", field)
}

type UserGoValidator struct{}

func (v UserGoValidator) Validate(data *entities.User) error {
	if !govalidator.IsUUIDv4(data.Id) {
		return ErrInvalidField("id")

	} else if !govalidator.IsEmail(data.Email) {
		return ErrInvalidField("email")

	} else if !govalidator.IsURL(data.AvatarUrl) {
		return ErrInvalidField("avatar_url")

	}

	return nil
}
