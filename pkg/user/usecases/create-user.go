package usecase

import (
	"context"
	"helpy/pkg/entities"
	"helpy/pkg/user"
	"helpy/pkg/utils"
	"time"
)

type CreateUser struct {
	Uuid      utils.UUIDGenerator
	Hasher    utils.Hasher
	Validator user.Validator
	// Repository user.Repository
}

func (usecase CreateUser) Create(ctx context.Context, data entities.User) (*entities.User, error) {
	uuid, err := usecase.Uuid.Generate()
	if err != nil {
		return nil, err
	}

	data.Id = uuid

	hash, err := usecase.Hasher.Hash(data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = hash

	err = usecase.Validator.Validate(&data)
	if err != nil {
		return nil, err
	}

	data.CreatedAt = time.Now().Local()
	data.UpdatedAt = time.Now().Local()

	// existing, err := usecase.Repository.Exists(ctx, data.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// if existing {
	// 	return nil, user.ErrExistingEmail
	// }

	// err = usecase.Repository.Create(ctx, &data)
	// if err != nil {
	// 	return nil, err
	// }

	return &data, nil
}
