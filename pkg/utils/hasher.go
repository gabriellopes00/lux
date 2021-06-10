package utils

type Hasher interface {
	Hash(payload *string) error
}
