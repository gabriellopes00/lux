package usecases

import (
	"context"
	"lux/pkg/entities"
	"lux/pkg/user"
	"lux/pkg/utils"
	"time"
)

type CreateUser struct {
	UuidGenerator utils.UUIDGenerator
	Hasher        utils.Hasher
	Repository    user.Repository
}

func (usecase CreateUser) Create(ctx context.Context, data entities.User) (*entities.User, error) {
	existing, err := usecase.Repository.Exists(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if existing {
		return nil, user.ErrExistingEmail
	}

	uuid, err := usecase.UuidGenerator.Generate()
	if err != nil {
		return nil, err
	}

	data.Id = uuid

	hash, err := usecase.Hasher.GenHash(data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = hash

	data.CreatedAt = time.Now().Local()
	data.UpdatedAt = time.Now().Local()

	err = usecase.Repository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
