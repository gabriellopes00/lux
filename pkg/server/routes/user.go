package router

import (
	build "helpy/pkg/server/build/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{Path: "/user", Method: http.MethodPost, Handler: build.CreateUserHandler.Handle},
}
