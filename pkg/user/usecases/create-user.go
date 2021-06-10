package usecase

import (
	"context"
	"helpy/pkg/entities"
	"helpy/pkg/user"
	"helpy/pkg/utils"
	"time"
)

type CreateUser struct {
	uuid       utils.UUIDGenerator
	hasher     utils.Hasher
	validator  user.Validator
	repository user.Repository
}

func (c *CreateUser) Create(ctx context.Context, data entities.User) (*entities.User, error) {
	uuid, err := c.uuid.Generate()
	if err != nil {
		return nil, user.InternalProcessingErr(err.Error())
	}
	data.Id = uuid

	err = c.hasher.Hash(&data.Password)
	if err != nil {
		return nil, user.InternalProcessingErr(err.Error())
	}

	err = c.validator.Validate(&data)
	if err != nil {
		return nil, user.InternalProcessingErr(err.Error())
	}

	data.CreatedAt = time.Now().Local()
	data.UpdatedAt = time.Now().Local()

	existing, err := c.repository.Exists(ctx, data.Email)
	if err != nil {
		return nil, user.InternalProcessingErr(err.Error())
	} else if existing {
		return nil, user.ExistingEmailErr(data.Email)
	}

	err = c.repository.Create(ctx, &data)
	if err != nil {
		return nil, user.InternalProcessingErr(err.Error())
	}

	return &data, nil
}
