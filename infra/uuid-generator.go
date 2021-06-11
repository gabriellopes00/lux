package infra

import "github.com/gofrs/uuid"

type UUIDGenerator struct{}

func (u UUIDGenerator) Generate() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return uuid.String(), err
	}

	return uuid.String(), nil
}
