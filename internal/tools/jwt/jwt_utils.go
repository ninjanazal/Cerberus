package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(p_usrId string, p_durationM uint) (string, error) {
	var expiration time.Time = time.Now().Add(time.Duration(p_durationM) * time.Minute)

	claims := Claims{
		UserID: p_usrId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	var tkn *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tkn.SignedString(jwtSecret)
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func ValidateJWT(p_token string) (*Claims, error) {
	claims := Claims{}

	token, err := jwt.ParseWithClaims(p_token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
