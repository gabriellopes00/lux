package usecase

import (
	"context"
	"helpy/pkg/entities"
	"helpy/pkg/user"
	"time"
)

type UpdateUser struct {
	Validator  user.Validator
	Repository user.Repository
}

func (c *UpdateUser) Update(ctx context.Context, data entities.User) (*entities.User, error) {
	existing, err := c.Repository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	} else if existing.Id == "" {
		return nil, user.ErrNonExistentEmail
	}

	err = c.Validator.Validate(&data)
	if err != nil {
		return nil, err
	}

	data.UpdatedAt = time.Now().Local()

	err = c.Repository.Update(ctx, data.Email, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
