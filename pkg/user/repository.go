package user

import (
	"context"
	"helpy/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	Exists(ctx context.Context, email string) (bool, error)
	FindAvailable(ctx context.Context) (*[]entities.User, error)
}
