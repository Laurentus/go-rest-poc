package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, options *options.ClientOptions) (*mongo.Client, error) {
	return mongo.Connect(ctx, options)
}
