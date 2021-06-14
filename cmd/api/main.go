package main

import (
	"context"
	"fmt"
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
		Name:        "Cleisom2",
		Email:       "clsiom@mail.com",
		Password:    "helpyapp00",
		IsAvailable: false,
		AvatarUrl:   "https://asdffff.png",
		Gender:      "M",
		BirthDate:   time.Date(2005, 4, 13, 0, 0, 0, 0, time.Local),
	}
	_, err = userUsecase.Create(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}

	findAvailableUsecase := usecase.FindAvailableUser{
		Repository: repositories.PgUserRepository{Db: database},
	}
	users, err := findAvailableUsecase.FindAvaliable(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range *users {
		fmt.Println(v)
	}

}
