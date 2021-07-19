package builds

import (
	"lux/infra/db/repositories"
	"lux/infra/utils"
	"lux/infra/validation"
	"lux/pkg/api/handlers"
	"lux/pkg/services/auth"
	"lux/pkg/user/usecases"

	"gorm.io/gorm"
)

func NewCreateUserHandler(conn *gorm.DB) *handlers.CreateUserHandler {
	return &handlers.CreateUserHandler{
		Validator: validation.UserGoValidator{},
		Auth: auth.UserAuth{
			UUIDGenerator: utils.UUIDGenerator{},
		},
		Usecase: usecases.CreateUser{
			UuidGenerator: utils.UUIDGenerator{},
			Hasher:        utils.Argon2Hasher{},
			Repository:    repositories.PgUserRepository{Db: conn},
		},
	}
}
