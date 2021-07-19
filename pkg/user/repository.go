package user

import (
	"context"
	"lux/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	Exists(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}
