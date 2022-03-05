package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	TokenString string `json:"token"`
	ExpiredAt   int64  `json:"expired_at"`
	IssuedAt    int64  `json:"issued_at"`
	Issuer      string `json:"issuer"`
	secretKey   string
}

func NewJWTService() JWTService {
	return &jwtService{
		Issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	expiredAt := time.Now().Add(time.Hour * 1).Unix()

	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
