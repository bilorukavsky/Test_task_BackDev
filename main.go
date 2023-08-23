package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestBody struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}

type RequestBodyLogin struct {
	Username string `json:"username"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора пользователя из параметров запроса
	var requestBody RequestBodyLogin
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	UserName := requestBody.Username
	fmt.Println(UserName)

	accessTokenString, err := generateAccessToken(UserName)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

	// Генерация Refresh токена и хеша для хранения
	refreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Invalid generate refresh token hash", http.StatusInternalServerError)
		return
	}

	// Создание хеша испульза refresh токен
	refreshTokenHash, err := hashToken(refreshToken)
	if err != nil {
		http.Error(w, "Could not generate refresh token hash", http.StatusInternalServerError)
		return
	}

	err = saveRefreshTokenHash(UserName, refreshTokenHash)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Отправка Access и Refresh токенов в ответе
	response := TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	// Получение Refresh токена и username из параметров запроса
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	UserName := requestBody.Username
	refreshToken := requestBody.RefreshToken

	dbToken, err := searchRefreshTokenHash(UserName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Invalid UserName", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err = compareHashAndToken(dbToken.Hash, refreshToken)
	if err != nil {
		http.Error(w, "Invalid RefreshToken", http.StatusInternalServerError)
		return
	}

	newAccessTokenString, err := generateAccessToken(UserName)
	if err != nil {
		http.Error(w, "Could not generate new access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Invalid generate refresh token hash", http.StatusInternalServerError)
		return
	}

	newRefreshTokenHash, err := hashToken(newRefreshToken)
	if err != nil {
		http.Error(w, "Could not generate refresh token hash", http.StatusInternalServerError)
		return
	}

	err = updateHashForUser(UserName, newRefreshTokenHash)
	if err != nil {
		http.Error(w, "Invalid update data base", http.StatusInternalServerError)
		return
	}

	response := TokenPair{
		AccessToken:  newAccessTokenString,
		RefreshToken: newRefreshTokenHash,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Eror loading .env file")
	}

	initDB()
	defer closeDB()

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/refresh", refreshHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
