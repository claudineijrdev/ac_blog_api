package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func NewDB(connectionString string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if db == nil {
		opts := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}
		db = client

	}
	err := db.Ping(ctx, nil)
	if err != nil {
		db = nil
		return NewDB(connectionString)
	}
	return db
}

func GetCollection(collectionName, database string) *mongo.Collection {
	return db.Database(database).Collection(collectionName)
}
