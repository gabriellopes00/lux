package tests

import (
	"lux/infra/utils"
	"testing"

	"github.com/asaskevich/govalidator"
)

func TestGenerate(t *testing.T) {
	uuidGenerator := utils.UUIDGenerator{}
	uuid, _ := uuidGenerator.Generate()

	if !govalidator.IsUUIDv4(uuid) {
		t.Errorf("generated uuid is not a valid uuid v4")
	}
}
