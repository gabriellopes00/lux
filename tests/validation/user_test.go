package tests

import (
	"lux/infra/validation"
	"lux/pkg/entities"
	"testing"
	"time"
)

func TestValdiate(t *testing.T) {
	validator := validation.UserGoValidator{}
	fakeUser := entities.User{
		Id:        "5aa30a03-3078-4418-adf6-64305b62be89",
		Name:      "User Name",
		Email:     "user@mail.com",
		Password:  "user_pass00",
		Bio:       "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Dicta natus voluptatem non delectus magnam aliquam aliquid, hic vitae architecto autem corporis perspiciatis eum qui placeat accusantium repellat deleniti et ipsam?",
		AvatarUrl: "https://user_avatar.png",
		BirthDate: time.Date(2005, 4, 13, 0, 0, 0, 0, time.Local),
	}

	t.Run("Avatar Url Validation", func(t *testing.T) {
		helper := fakeUser
		helper.AvatarUrl = "invalid_url"
		err := validator.Validate(&helper)
		if err == nil {
			t.Errorf("avatar url validation failure")
		}

		helper.AvatarUrl = ""
		err = validator.Validate(&helper)
		if err == nil {
			t.Errorf("avatar url validation failure")
		}
	})

	t.Run("Email Validation", func(t *testing.T) {
		helper := fakeUser
		helper.Email = "invalidmail.com"
		err := validator.Validate(&helper)
		if err == nil {
			t.Errorf("email validation failure")
		}

		helper.Email = "invalid@mail"
		err = validator.Validate(&helper)
		if err == nil {
			t.Errorf("email validation failure 2")
		}

		helper.Email = ""
		err = validator.Validate(&helper)
		if err == nil {
			t.Errorf("email validation failure")
		}

		helper.Email = "user@mail.com"
		err = validator.Validate(&helper)
		if err != nil {
			t.Errorf("email validation failure 3")
		}
	})
}
