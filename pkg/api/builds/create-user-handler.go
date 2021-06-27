package builds

import (
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	"helpy/infra/validation"
	h "helpy/pkg/api/handlers"
	usecase "helpy/pkg/user/usecases"

	"gorm.io/gorm"
)

func NewCreateUserHandler(conn *gorm.DB) *h.CreateUserHandler {
	return &h.CreateUserHandler{
		Usecase: usecase.CreateUser{
			Uuid:      utils.UUIDGenerator{},
			Hasher:    utils.Argon2Hasher{},
			Validator: validation.UserGoValidator{},
			Repository: repositories.PgUserRepository{
				Db: conn,
			},
		},
	}
}
