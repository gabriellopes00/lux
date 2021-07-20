package article

import (
	"context"
	"lux/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.Article) error
	FindById(ctx context.Context, email string) (*entities.Article, error)
}
