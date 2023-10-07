package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Ctx context.Context
var Cancel context.CancelFunc
var Db *mongo.Database

func Connect() *mongo.Client {
	uri, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		log.Fatal("Must set mongo uri as env variable: MONGO_URI")
	}

	dbName, ok := os.LookupEnv("MONGO_DB_NAME")
	if !ok {
		log.Fatal("Must set mongo db name as env variable: MONGO_DB_NAME")
	}

	log.Println("Setting up mongo...")

	Ctx, Cancel = context.WithCancel(context.Background())

	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	Db = client.Database(dbName)
	Client = client

	return client
}

func Close() {
	if err := Client.Disconnect(Ctx); err != nil{
		panic(err)
	}
	Cancel()
	log.Println("Mongo Closed")
}