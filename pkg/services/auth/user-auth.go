package auth

import (
	"lux/config/env"
	"lux/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserAuthentication interface {
	GenAuthToken(payload interface{}) (string, error)
}

type UserAuth struct {
	UUIDGenerator utils.UUIDGenerator
}

func (a *UserAuth) GenAuthToken(payload interface{}) (string, error) {
	claims := jwt.MapClaims{}

	id, err := a.UUIDGenerator.Generate()
	if err != nil {
		return "", err
	}

	claims["authorized"] = true
	claims["userId"] = payload
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(env.TOKEN_EXPIRATION_TIME)).Local()

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.TOKEN_PRIVATE_KEY))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signed, nil
}

// func ValidateAuthToken(token string) {

// 	verified, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected sign method %v", t.Header["alg"])
// 		}

// 		return jwt.ParseRSAPublicKeyFromPEM([]byte(env.TOKEN_PUBLIC_KEY))
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("")
// 	}

// 	fmt.Println(verified)
// 	fmt.Println("")

// 	claims := verified.Claims.(jwt.MapClaims)

// 	fmt.Println(claims["user_id"])
// 	fmt.Println("")

// }
