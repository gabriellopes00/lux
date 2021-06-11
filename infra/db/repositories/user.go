package repositories

import (
	"context"
	"database/sql"
	"helpy/pkg/entities"
	"log"

	_ "github.com/lib/pq"
)

type PgUserRepository struct {
	Db *sql.DB
}

func (r PgUserRepository) Create(ctx context.Context, user *entities.User) error {
	stm, err := r.Db.Prepare("INSERT INTO \"user\" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		user.Id,
		user.Name,
		user.Email,
		user.Password,
		user.IsAvailable,
		user.AvatarUrl,
		user.Gender,
		user.BirthDate,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	)

	if err != nil {
		return err
	}

	return nil
	// fmt.Printf("name: %s\n", user.Name)
	// fmt.Printf("birthDate: %s\n", user.BirthDate)
	// fmt.Printf("id: %s\n", user.Id)
	// fmt.Printf("email: %s\n", user.Email)
	// fmt.Printf("avatarUrl: %s\n", user.AvatarUrl)
	// fmt.Printf("createdAt: %s\n", user.CreatedAt)
	// fmt.Printf("updatedAt: %s\n", user.UpdatedAt)
	// fmt.Printf("hashPass: %s\n", user.Password)
	// fmt.Printf("gender: %s\n", user.Gender)
	// fmt.Printf("deletedAT: %s\n", user.DeletedAt)
}

func (r PgUserRepository) Exists(ctx context.Context, email string) (bool, error) {
	return false, nil
}
