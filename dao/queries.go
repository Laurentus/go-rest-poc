package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DAO interface {
	All() chan interface{}
}

func FindAll(collection *mongo.Collection, handler func(*mongo.Cursor), closer func()) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Always close connection
	defer cur.Close(ctx)

	// On demand fetching of campaigns
	for cur.Next(ctx) {
		handler(cur)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	closer()
}
