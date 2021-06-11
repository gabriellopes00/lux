package main

import (
	"context"
	"helpy/infra"
	postgres "helpy/infra/repositories/pg"
	valid "helpy/infra/validation"
	"helpy/pkg/entities"
	usecase "helpy/pkg/user/usecases"
	"time"
)

func main() {
	userUsecase := usecase.CreateUser{
		Uuid:       infra.UUIDGenerator{},
		Hasher:     infra.Argon2Hasher{},
		Repository: postgres.PgUserRepository{},
		Validator:  valid.UserGoValidator{},
	}

	user := entities.User{
		Name:        "Gabriel Lopes",
		Email:       "gabriel@mail.com",
		Password:    "helpyapp00",
		IsAvailable: true,
		AvatarUrl:   "https://avatar.png",
		Gender:      "M",
		BirthDate:   time.Date(2004, 4, 13, 12, 30, 0, 0, time.Local),
	}
	ctx := context.Background()
	userUsecase.Create(ctx, user)
}
