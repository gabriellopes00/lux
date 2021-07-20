package usecases

import (
	"context"
	"lux/pkg/article"
	"lux/pkg/entities"
	"lux/pkg/utils"
	"time"
)

type CreateArticle struct {
	UuidGenerator utils.UUIDGenerator
	Repository    article.Repository
}

func (usecase CreateArticle) Create(ctx context.Context, data entities.Article) (*entities.Article, error) {
	uuid, err := usecase.UuidGenerator.Generate()
	if err != nil {
		return nil, err
	}

	data.Id = uuid

	data.CreatedAt = time.Now().Local()
	data.UpdatedAt = time.Now().Local()

	err = usecase.Repository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
