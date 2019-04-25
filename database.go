package auth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	fmt.Println(uri)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Database connection error  %s", err)
	}

	ctx, _ = context.WithTimeout(context.TODO(), 2*time.Second)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Database connection error  %s", err)
	}
	log.Print("Database connection is establishing.")
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
