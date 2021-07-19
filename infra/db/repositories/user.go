package repositories

import (
	"context"
	"lux/pkg/entities"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type PgUserRepository struct {
	Db *gorm.DB
}

func (repository PgUserRepository) Create(ctx context.Context, user *entities.User) error {
	err := repository.Db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository PgUserRepository) Exists(ctx context.Context, email string) (bool, error) {
	user := &entities.User{}

	result := repository.Db.Find(user, "email = ?", email)
	if result.Error != nil {
		return false, result.Error
	}

	if user.Email == email {
		return true, nil
	}

	return false, nil
}

func (repository PgUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := entities.User{}

	result := repository.Db.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
