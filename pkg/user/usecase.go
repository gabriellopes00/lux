package user

import (
	"context"
	"lux/pkg/entities"
)

type CreateUser interface {
	Create(ctx context.Context, data entities.User) (*entities.User, error)
}
