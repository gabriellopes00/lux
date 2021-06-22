package router

import (
	"helpy/infra/utils"
	"helpy/infra/validation"
	"helpy/pkg/server/handlers"
	usecase "helpy/pkg/user/usecases"
	"net/http"
)

var CreateUserHandler = handlers.CreateUserHandler{
	Usecase: usecase.CreateUser{
		Uuid:      utils.UUIDGenerator{},
		Hasher:    utils.Argon2Hasher{},
		Validator: validation.UserGoValidator{},
	},
}

var UserRoutes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: CreateUserHandler.Handle},
}
