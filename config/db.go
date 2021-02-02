package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mongodb consts
const (
	MongoDefaultTimeout = 2 // default timeout
)

type DBConnection struct {
	client *mongo.Client
}

func NewConnection(host string) (client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), MongoDefaultTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}
