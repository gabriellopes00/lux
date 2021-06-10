package usecase

import "helpy/pkg/usecases/user"

type CreateUser struct {
	repository user.Repository
	validator  user.Validator
}
