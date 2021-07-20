package article

import (
	"context"
	"lux/pkg/entities"
)

type CreateArticle interface {
	Create(ctx context.Context, data entities.Article) (*entities.Article, error)
}
