package utils

type UUIDGenerator interface {
	Generate() (string, error)
}
