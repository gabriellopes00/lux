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
	orm := db.GormPG{}
	database, err := orm.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	userUsecase := usecase.CreateUser{
		Uuid:       utils.UUIDGenerator{},
		Hasher:     utils.Argon2Hasher{},
		Repository: repositories.PgUserRepository{Db: database},
		Validator:  valid.UserGoValidator{},
	}

	user := entities.User{
		Name:        "Gabriel Lopes",
		Email:       "gabriel@mail.com",
		Password:    "helpyapp00",
		IsAvailable: true,
		AvatarUrl:   "https://avatar.png",
		Gender:      "M",
		BirthDate:   time.Date(2005, 4, 13, 0, 0, 0, 0, time.Local),
	}
	userUsecase.Create(context.Background(), user)
}
