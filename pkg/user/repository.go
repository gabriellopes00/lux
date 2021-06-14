package user

import (
	"context"
	"helpy/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, email string, data *entities.User) error
	Exists(ctx context.Context, email string) (bool, error)
	FindAvailable(ctx context.Context) (*[]entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}
