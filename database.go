package auth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database struct {
	URL            string
	Username       string
	Password       string
	DatabaseName   string
	CollectionName string
	Client         mongo.Client
	Collection     mongo.Collection
}

// ConnectDatabase is takes username, password, url and gives database client
func ConnectDatabase(username, password, url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s",
			username, password, url)))

	if err != nil {
		log.Fatalf("Connection error  %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Connection error  %v", err)
	}
	return client
}

// InitDatabase takes parameters from environment and creates a Database Object
func InitDatabase(url, dbName, username, password, collectionName string) *Database {
	db := new(Database)
	db.Username = username
	db.Password = password
	db.URL = url
	db.DatabaseName = dbName
	db.CollectionName = collectionName
	db.Client = *ConnectDatabase(username, password, url)
	db.Collection = *db.Client.Database(db.DatabaseName).Collection(db.CollectionName)
	return db
}
