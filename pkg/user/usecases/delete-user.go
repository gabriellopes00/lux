package usecases

import (
	"context"
	"helpy/pkg/user"
)

type DeleteUser struct {
	Repository user.Repository
}

func (usecase DeleteUser) Delete(ctx context.Context, id string) error {
	return usecase.Repository.Delete(ctx, id)
}
