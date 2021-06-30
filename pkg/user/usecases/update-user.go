package usecases

import (
	"helpy/pkg/user"
)

type UpdateUser struct {
	Repository user.Repository
}

// func (usecase *UpdateUser) Update(ctx context.Context, data entities.User) (*entities.User, error) {
// 	existing, err := usecase.Repository.FindByEmail(ctx, data.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if existing == nil {
// 		return nil, user.ErrNonExistentEmail
// 	}

// 	if err = usecase.Validator.Validate(&data); err != nil {
// 		return nil, err
// 	}

// 	data.UpdatedAt = time.Now().Local()

// 	if err = usecase.Repository.Update(ctx, &data); err != nil {
// 		return nil, err
// 	}

// 	return &data, nil
// }
