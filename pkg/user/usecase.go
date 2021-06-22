package user

import (
	"context"
	"helpy/pkg/entities"
)

type CreateUser interface {
	Create(ctx context.Context, data entities.User) (*entities.User, error)
}

type FindAvailableUser interface {
	Find(ctx context.Context) (*[]entities.User, error)
}

type DeleteUser interface {
	Delete(ctx context.Context, id string) error
}
