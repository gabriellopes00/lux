package usecase

import (
	"context"
	"helpy/pkg/entities"
	"helpy/pkg/user"
	"helpy/pkg/utils"
	"time"
)

type CreateUser struct {
	Uuid       utils.UUIDGenerator
	Hasher     utils.Hasher
	Validator  user.Validator
	Repository user.Repository
}

func (c *CreateUser) Create(ctx context.Context, data entities.User) (*entities.User, error) {
	uuid, err := c.Uuid.Generate()
	if err != nil {
		return nil, err
	}
	data.Id = uuid

	hash, err := c.Hasher.Hash(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = hash

	err = c.Validator.Validate(&data)
	if err != nil {
		return nil, err
	}

	data.CreatedAt = time.Now().Local()
	data.UpdatedAt = time.Now().Local()

	existing, err := c.Repository.Exists(ctx, data.Email)
	if err != nil {
		return nil, err
	} else if existing {
		return nil, user.ErrExistingEmail
	}

	err = c.Repository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
