package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB(uri string) *mongo.Client {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(tx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	//TODO:What is this for?
	// MongoClient = client
	return client

}

func GetCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
