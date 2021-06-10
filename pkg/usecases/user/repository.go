package user

import (
	"context"
	"helpy/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
}
