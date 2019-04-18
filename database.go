package auth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// Environment variables with keys
const ENV_USERNAME = "DB_USERNAME"
const ENV_PASSWORD = "DB_PASSWORD"
const ENV_URL = "DB_URL"
const ENV_DB_NAME = "DB_NAME"
const ENV_COLLECTION = "DB_COLLECTION"

type Database struct {
	Username     string
	Password     string
	URL          string
	DatabaseName string
	Collection   string
	Client       mongo.Client
}

// ConnectDatabase is takes username, password, url and gives database client
func ConnectDatabase(username, password, url string) (*mongo.Client) {
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

// Get Collection takes database and collection name and gives a MongoDB Collection
func GetCollection(database, collection string, client *mongo.Client) *mongo.Collection {
	return client.Database("auth").Collection("users")
}

// GetDBClient takes all database data from environment for opening connection
// and choosing collection, gives a usable DBClient
func GetDBClient(username, password, url, database, collection string) *mongo.Collection {
	client := ConnectDatabase("go-auth", "IOvrErajOi11TY7q",
		"graphql-backend-mongo-cluster-shard-00-00-pej2l.mongodb.net:27017,graphql-backend-mongo-cluster-shard-00-01-pej2l.mongodb.net:27017,graphql-backend-mongo-cluster-shard-00-02-pej2l.mongodb.net:27017/test?ssl=true&replicaSet=graphql-backend-mongo-cluster-shard-0&authSource=admin&retryWrites=true")
	return client.Database("auth").Collection("users")
}

// InitDatabase takes parameters from environment and creates a Database Object
func InitDatabase() {
	username := os.Getenv(ENV_USERNAME)
	password := os.Getenv(ENV_PASSWORD)
	url := os.Getenv(ENV_URL)
	dbName := os.Getenv(ENV_DB_NAME)
	collection := os.Getenv(ENV_COLLECTION)

	if username == "" {
		logFatal(ENV_USERNAME)
	}

	if password == "" {
		logFatal(ENV_PASSWORD)
	}
	// ...
	db := new(Database)
	db.Username = username
	db.Password = password
	db.URL = url
	db.DatabaseName = dbName
	db.Collection = collection

}

func logFatal(key string) {
	log.Fatalf("%s is required!", key)
}
