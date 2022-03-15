package service

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateJWT(name string, admin bool) string
	ValidateJWT(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	expiresAt int64
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "secret",
		expiresAt: time.Now().AddDate(1, 0, 0).Unix(),
	}
}

func (s *jwtService) GenerateJWT(name string, admin bool) string {
	claims := jwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: s.expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return ts
}

func (s *jwtService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
}
