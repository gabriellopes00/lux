package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"helpy/pkg/api/validators"
	"helpy/pkg/entities"
	"helpy/pkg/services/auth"
	u "helpy/pkg/user"
	"io"
	"net/http"
)

type CreateUserHandler struct {
	Validator validators.UserValidator
	Usecase   u.CreateUser
	Auth      auth.UserAuth

	response struct {
		AccessToken string        `json:"access_token"`
		User        entities.User `json:"user"`
	}
}

func (handler *CreateUserHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	var user entities.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	err = handler.Validator.Validate(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	created, err := handler.Usecase.Create(context.Background(), user)
	if errors.Is(err, u.ErrExistingEmail) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return

	} else if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := handler.Auth.GenAuthToken(created.Id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	handler.response.AccessToken = token
	handler.response.User = *created

	err = json.NewEncoder(rw).Encode(handler.response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
