package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Генерации refresh токенов
func generateRefreshToken() (string, error) {
	tokenLength := 32 // Длина Refresh токена в байтах
	tokenBytes := make([]byte, tokenLength)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(tokenBytes), nil
}

// Генерации хеша из refresh токенов
func hashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Проверка хэша
func compareHashAndToken(refreshTokenHash string, refreshToken string) error {
	err := bcrypt.CompareHashAndPassword([]byte(refreshTokenHash), []byte(refreshToken))
	return err
}

// Генерация AccessToken
func generateAccessToken(username string) (string, error) {
	accessClaims := jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}
