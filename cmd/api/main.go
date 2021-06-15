package main

import (
	"context"
	"fmt"
	"helpy/config"
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
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("error while loading environment variables: \n %s", err.Error())
	}

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
		Email:       "gabriellopess@mail.com",
		Password:    "helpyapp00",
		IsAvailable: true,
		AvatarUrl:   "https://avatar.png",
		Gender:      "M",
		BirthDate:   time.Date(2005, 4, 13, 0, 0, 0, 0, time.Local),
	}
	created, err := userUsecase.Create(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(created)

	// findAvailableUsecase := usecase.FindAvailableUser{
	// 	Repository: repositories.PgUserRepository{Db: database},
	// }
	// users, err := findAvailableUsecase.FindAvaliable(context.Background())
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, v := range *users {
	// 	fmt.Println(v)
	// }

	// asdf := usecase.DeleteUser{
	// 	Repository: repositories.PgUserRepository{Db: database},
	// }

	// err = asdf.Delete(context.Background(), created.Id)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	time.Sleep(time.Second * 30)

	asdf := usecase.UpdateUser{
		Repository: repositories.PgUserRepository{Db: database},
		Validator:  valid.UserGoValidator{},
	}

	created.Name = "asdf"
	created.AvatarUrl = "http://asdf.jpg"

	asdf.Update(context.Background(), *created)

}
