package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func StartDB() {
	// Database Config
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic("Failed whiling mongo.NewClient")
	}

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic("Failed Context to mongodb")
	}

	//Cancel context to avoid memory leak
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Database Connected!")
	}

	db = client.Database(os.Getenv("DB_NAME"))
}

func GetDatabase() *mongo.Database {
	return db
}

func UsersCollection() *mongo.Collection {
	user_collection := db.Collection("users")

	return user_collection
}
