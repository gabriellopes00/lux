package utils

type Hasher interface {
	GenHash(payload string) (string, error)
}
