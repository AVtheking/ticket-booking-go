package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claim struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashPassword)
}

func CheckPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func GenerateToken(id uint, email string) (string, error) {
	claims := &Claim{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
