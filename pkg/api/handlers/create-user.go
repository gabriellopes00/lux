package handlers

import (
	"context"
	"encoding/json"
	"helpy/pkg/entities"
	"helpy/pkg/user"
	"io"
	"net/http"
)

type CreateUserHandler struct {
	Validator user.Validator
	Usecase   user.CreateUser
}

func (h *CreateUserHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	var user entities.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		rw.Write([]byte(err.Error()))
		return
	}

	err = h.Validator.Validate(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	created, err := h.Usecase.Create(context.Background(), user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(rw).Encode(created)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
