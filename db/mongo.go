package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

type singletonClient *mongo.Client

var client singletonClient

func GetClient() *mongo.Client {
	once.Do(func() {
		client = connect()
	})
	return client
}

func GetDb() *mongo.Database {
	return GetClient().Database("app")
}

func connect() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongo, err := NewMongoClient(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = mongo.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return mongo
}
