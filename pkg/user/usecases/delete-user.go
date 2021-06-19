package usecase

import (
	"context"
	"helpy/pkg/user"
)

type DeleteUser struct {
	Repository user.Repository
}

func (usecase DeleteUser) Delete(ctx context.Context, id string) error {
	err := usecase.Repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
