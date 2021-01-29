package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_TIMEOUT_SEC = 2

type DBConnection struct {
	client *mongo.Client
}

func NewConnection(host string) (client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		panic(err)
	}
	return client
}
