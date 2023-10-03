package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go/v4"
)

type IJWTService interface {
	GenerateToken(userId int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error)
}

type JWTService struct{}

func NewServiceJWT() *JWTService {
	return &JWTService{}
}

const (
	SECRET_KEY = "fjbgehfbgueryguefygru"
)

func (s *JWTService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *JWTService) ValidasiToken(token string) (*jwt.Token, error) {
	myToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return myToken, err
	}

	return myToken, nil
}
