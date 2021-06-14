package repositories

import (
	"context"
	"helpy/pkg/entities"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type PgUserRepository struct {
	Db *gorm.DB
}

func (repository PgUserRepository) Create(ctx context.Context, user *entities.User) error {
	result := repository.Db.Create(user)
	if result.Error != nil {
		return result.Error
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

func (repository PgUserRepository) FindAvailable(ctx context.Context) (*[]entities.User, error) {
	users := &[]entities.User{}
	result := repository.Db.Find(users, "is_available = ?", true)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
