package repositories

import (
	"context"
	"lux/pkg/entities"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type PgArticleRepository struct {
	Db *gorm.DB
}

func (repository PgArticleRepository) Create(ctx context.Context, article *entities.Article) error {
	return repository.Db.Create(article).Error
}

func (repository PgArticleRepository) FindById(ctx context.Context, id string) (*entities.Article, error) {
	article := entities.Article{}

	result := repository.Db.Find(&article, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}
