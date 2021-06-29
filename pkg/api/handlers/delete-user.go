package handlers

import (
	"context"
	"errors"
	"helpy/pkg/user"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type DeleteUserHandler struct {
	Usecase user.DeleteUser
}

func (h *DeleteUserHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	if !govalidator.IsUUIDv4(id) {
		err := errors.New("invalid user id param")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	err := h.Usecase.Delete(context.Background(), id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.WriteHeader(http.StatusOK)
}
