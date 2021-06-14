package user

import (
	"context"
	"helpy/pkg/entities"
)

type Usecase interface {
	Create(ctx context.Context, data entities.User) (*entities.User, error)
	FindAvaliable(ctx context.Context) (*[]entities.User, error)
}
