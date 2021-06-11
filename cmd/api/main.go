package main

import (
	"context"
	"helpy/infra/db"
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	valid "helpy/infra/validation"
	"helpy/pkg/entities"
	usecase "helpy/pkg/user/usecases"
	"log"
	"time"
)

func main() {
	database, err := db.ConnectPg()
	if err != nil {
		log.Fatalln(err)
	}
	userUsecase := usecase.CreateUser{
		Uuid:   utils.UUIDGenerator{},
		Hasher: utils.Argon2Hasher{},
		Repository: repositories.PgUserRepository{
			Db: database,
		},
		Validator: valid.UserGoValidator{},
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
