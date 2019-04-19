package auth

import (
	"log"
	"os"
)

/*
username: go-auth
password: IOvrErajOi11TY7q
url: graphql-backend-mongo-cluster-shard-00-00-pej2l.mongodb.net:27017,graphql-backend-mongo-cluster-shard-00-01-pej2l.mongodb.net:27017,graphql-backend-mongo-cluster-shard-00-02-pej2l.mongodb.net:27017/test?ssl=true&replicaSet=graphql-backend-mongo-cluster-shard-0&authSource=admin&retryWrites=true
dbName: auth
collection: users
*/

// Keys of required Environment variables
const (
	ENV_DB_URL        = "DB_URL"
	ENV_DB_NAME       = "DB_NAME"
	ENV_DB_USERNAME   = "DB_USERNAME"
	ENV_DB_PASSWORD   = "DB_PASSWORD"
	ENV_DB_COLLECTION = "DB_COLLECTION"
)

// Environment type is specify the required parameters for startup
type Environment struct {
	DB_URL        string
	DB_NAME       string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_COLLECTION string
}

// GetEnv func takes all needed environment variables fro OS
func GetEnv() *Environment {
	username := os.Getenv(ENV_DB_USERNAME)
	password := os.Getenv(ENV_DB_PASSWORD)
	url := os.Getenv(ENV_DB_URL)
	dbName := os.Getenv(ENV_DB_NAME)
	collection := os.Getenv(ENV_DB_COLLECTION)

	if url == "" {
		log.Fatalf("%s is required!", ENV_DB_URL)
	}
	if dbName == "" {
		log.Fatalf("%s is required!", ENV_DB_NAME)
	}
	if collection == "" {
		log.Fatalf("%s is required!", ENV_DB_COLLECTION)
	}
	env := new(Environment)
	env.DB_URL = url
	env.DB_NAME = dbName
	env.DB_USERNAME = username
	env.DB_PASSWORD = password
	env.DB_COLLECTION = collection
	return env
}
