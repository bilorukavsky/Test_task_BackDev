package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type RefreshTokenHashBD struct {
	Username string `bson:"username"`
	Hash     string `bson:"hash"`
}

func initDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("HOST"))
	var err error

	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func closeDB() {
	if client != nil {
		_ = client.Disconnect(context.Background())
	}
}

func createTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// Функция для сохранения хеша и username в бд
func saveRefreshTokenHash(username, refreshTokenHash string) error {
	ctx, cancel := createTimeoutContext()
	defer cancel()

	collection := client.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION"))

	_, err := collection.InsertOne(ctx, RefreshTokenHashBD{Username: username, Hash: refreshTokenHash})
	if err != nil {
		return err
	}
	return nil
}

func searchRefreshTokenHash(username string) (RefreshTokenHashBD, error) {
	ctx, cancel := createTimeoutContext()
	defer cancel()

	collection := client.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION"))

	// Поиск токена в базе данных
	var dbToken RefreshTokenHashBD
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&dbToken)
	return dbToken, err
}

func updateHashForUser(username, newHash string) error {
	ctx, cancel := createTimeoutContext()
	defer cancel()

	collection := client.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION"))

	// Определение операции обновления
	update := bson.M{"$set": bson.M{"hash": newHash}}

	// Выполнение операции обновления
	_, err := collection.UpdateOne(ctx, bson.M{"username": username}, update)
	if err != nil {
		return err
	}

	return nil
}
