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

func CreateURIWithCredentials(username, password, url string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s",
		username, password, url)
}

func CreateURI(url string) string {
	return fmt.Sprintf("mongodb://%s", url)
}

// ConnectDatabase is takes username, password, url and gives database client
func ConnectDatabase(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

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

// Initialize takes parameters from environment and creates a Database Object
func Initialize(url, dbName, username, password, collectionName string) *Database {
	db := new(Database)
	db.URL = url
	db.DatabaseName = dbName
	db.CollectionName = collectionName
	if username != "" || password != "" {
		db.Username = username
		db.Password = password
		db.Client = *ConnectDatabase(CreateURIWithCredentials(username, password, url))
	} else {
		db.Client = *ConnectDatabase(CreateURI(url))
	}
	db.Collection = *db.Client.Database(db.DatabaseName).Collection(db.CollectionName)
	return db
}
