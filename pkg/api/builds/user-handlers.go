package builds

import (
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	"helpy/infra/validation"
	"helpy/pkg/api/handlers"
	"helpy/pkg/services/auth"
	"helpy/pkg/user/usecases"

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

func NewDeleteUserHandler(conn *gorm.DB) *handlers.DeleteUserHandler {
	return &handlers.DeleteUserHandler{
		Usecase: usecases.DeleteUser{
			Repository: repositories.PgUserRepository{Db: conn},
		},
	}
}
