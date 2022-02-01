package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTCustomClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func newJWTService() JWTService {
	return &jwtService{
		issuer:    "evan",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := "evansecret"
	return secretKey
}

func (s *jwtService) GenerateToken(UserId string) string {
	claims := &JWTCustomClaim{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		panic(err)
	}

	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token,
		func(t_ *jwt.Token) (interface{}, error) {
			if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sign method", t_.Header["alg"])
			}
			return []byte(s.secretKey), nil
		})
}
