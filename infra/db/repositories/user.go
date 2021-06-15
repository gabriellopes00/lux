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

	if err := repository.Db.Create(user).Error; err != nil {
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

func (repository PgUserRepository) FindAvailable(ctx context.Context) (*[]entities.User, error) {
	users := &[]entities.User{}
	result := repository.Db.Find(users, "is_available = ?", true)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository PgUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	users := &entities.User{}
	result := repository.Db.Find(users, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository PgUserRepository) Delete(ctx context.Context, id string) error {
	if err := repository.Db.Delete(&entities.User{Id: id}).Error; err != nil {
		return err
	}

	return nil
}

func (repository PgUserRepository) Update(ctx context.Context, data *entities.User) error {
	if err := repository.Db.Save(data).Error; err != nil {
		return err
	}

	return nil
}
