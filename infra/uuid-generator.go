package main

import (
	"helpy/pkg/utils"

	"github.com/gofrs/uuid"
)

type UUIDGenerator struct{}

func (u *UUIDGenerator) Generate() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return uuid.String(), utils.InternalProcessingErr(err.Error())
	}

	return uuid.String(), nil
}
