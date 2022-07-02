package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var ConnectURL = "mongodb://root:root@localhost:27017"

func ConnectMongo() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(ConnectURL))

	return client, ctx, cancel, err

}

func CloseMongo(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
