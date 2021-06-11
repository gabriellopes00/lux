package postgres

import (
	"context"
	"fmt"
	"helpy/pkg/entities"
)

type PgUserRepository struct{}

func (r PgUserRepository) Create(ctx context.Context, user *entities.User) error {
	fmt.Printf("name: %s\n", user.Name)
	fmt.Printf("birthDate: %s\n", user.BirthDate)
	fmt.Printf("id: %s\n", user.Id)
	fmt.Printf("email: %s\n", user.Email)
	fmt.Printf("avatarUrl: %s\n", user.AvatarUrl)
	fmt.Printf("createdAt: %s\n", user.CreatedAt)
	fmt.Printf("updatedAt: %s\n", user.UpdatedAt)
	fmt.Printf("hashPass: %s\n", user.Password)
	fmt.Printf("gender: %s\n", user.Gender)
	fmt.Printf("deletedAT: %s\n", user.DeletedAt)
	return nil
}

func (r PgUserRepository) Exists(ctx context.Context, email string) (bool, error) {
	return false, nil
}
