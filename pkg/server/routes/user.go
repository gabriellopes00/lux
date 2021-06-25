package router

import (
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	"helpy/infra/validation"
	"helpy/pkg/server/handlers"
	usecase "helpy/pkg/user/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

var CreateUserHandler = handlers.CreateUserHandler{
	Usecase: usecase.CreateUser{
		Uuid:       utils.UUIDGenerator{},
		Hasher:     utils.Argon2Hasher{},
		Validator:  validation.UserGoValidator{},
		Repository: repositories.PgUserRepository{},
	},
}

func SetupUserRoutes(r *mux.Router) {
	r.HandleFunc("/user", CreateUserHandler.Handle).Methods(http.MethodPost)

}
