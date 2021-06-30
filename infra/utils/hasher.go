package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("argon2: hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("argon2: incompatible version of argon2 hash")
)

type params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var defaultParams = &params{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

type Argon2Hasher struct{}

func (a Argon2Hasher) GenHash(payload string) (hash string, err error) {
	salt, err := generateRandomBytes(defaultParams.SaltLength)
	if err != nil {
		return "", err
	}

	key := argon2.IDKey([]byte(payload), salt, defaultParams.Iterations, defaultParams.Memory, defaultParams.Parallelism, defaultParams.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Key := base64.RawStdEncoding.EncodeToString(key)

	hash = fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultParams.Memory,
		defaultParams.Iterations,
		defaultParams.Parallelism,
		b64Salt,
		b64Key,
	)
	return hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
