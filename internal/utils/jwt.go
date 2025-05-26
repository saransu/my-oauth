package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

type UserClaims struct {
	ID   uint
	Name string
	Age  int
	jwt.RegisteredClaims
}

func GenerateToken(u User) (string, error) {
	key, err := os.ReadFile("private.pem")
	if err != nil {
		return "", err
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}

	claims := UserClaims{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	return token.SignedString(privKey)
}

func DecryptToken(token string) (*UserClaims, error) {
	tk, err := jwt.ParseWithClaims(
		token,
		&UserClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodRS256 {
				return nil, errors.New("invalid signing method")
			}

			key, err := os.ReadFile("public.pem")
			if err != nil {
				return nil, err
			}

			pubKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
			if err != nil {
				return nil, err
			}

			return pubKey, nil
		},
	)

	if err != nil {
		return &UserClaims{}, err
	}

	u, ok := tk.Claims.(*UserClaims)
	if !ok {
		return &UserClaims{}, errors.New("invalid claims")
	}

	return u, err
}
