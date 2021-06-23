package build

import (
	"helpy/infra/db"
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	"helpy/infra/validation"
	"helpy/pkg/server/handlers"
	usecase "helpy/pkg/user/usecases"
	"log"

	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	orm := db.GormPG{}
	db, err := orm.Connect()
	if err != nil {
		log.Panicln(err)
	}

	err = orm.AutoMigrate(db)
	if err != nil {
		log.Panicln(err)
	}

	return db
}

var CreateUserHandler = handlers.CreateUserHandler{
	Usecase: usecase.CreateUser{
		Uuid:      utils.UUIDGenerator{},
		Hasher:    utils.Argon2Hasher{},
		Validator: validation.UserGoValidator{},
		Repository: repositories.PgUserRepository{
			Db: connectDb(),
		},
	},
}
