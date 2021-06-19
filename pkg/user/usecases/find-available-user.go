package usecase

import (
	"context"
	"helpy/pkg/entities"
	"helpy/pkg/user"
)

type FindAvailableUser struct {
	Repository user.Repository
}

func (usecase FindAvailableUser) FindAvaliable(ctx context.Context) (*[]entities.User, error) {
	users, err := usecase.Repository.FindAvailable(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
