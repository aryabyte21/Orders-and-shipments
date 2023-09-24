package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ConnectionString = "mongodb+srv://aryabyte:mkh4yt5fn4@cluster0.w0gftbh.mongodb.net/?retryWrites=true&w=majority"
	DBName = "stripedb"
)

func Init() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client.Database(DBName), nil
}

func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
