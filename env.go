package auth

import (
	"log"
	"os"
	"strconv"
	"time"
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
	EnvDbUrl        = "DB_URL"
	EnvDbName       = "DB_NAME"
	EnvDbUsername   = "DB_USERNAME"
	EnvDbPassword   = "DB_PASSWORD"
	EnvDbCollection = "DB_COLLECTION"
	EnvJwtSecret    = "JWT_SECRET"
	EnvJwtExpire    = "JWT_EXPIRE"
)

// Environment type is specify the required parameters for startup
type Environment struct {
	DbUrl        string
	DbName       string
	DbUsername   string
	DbPassword   string
	DbCollection string
	JwtSecret    string
	JwtExpire    time.Duration
}

var E *Environment

// GetEnv func takes all needed environment variables fro OS
func GetEnv() *Environment {
	username := os.Getenv(EnvDbUsername)
	password := os.Getenv(EnvDbPassword)
	url := os.Getenv(EnvDbUrl)
	dbName := os.Getenv(EnvDbName)
	collection := os.Getenv(EnvDbCollection)
	jwtSecret := os.Getenv(EnvJwtSecret)
	jwtExpire := os.Getenv(EnvJwtExpire)

	env := new(Environment)
	if jwtExpire != "" {
		exp, err := strconv.Atoi(jwtExpire)
		if err != nil {
			log.Fatalf("%s must be int!", EnvJwtExpire)
		}
		env.JwtExpire = time.Duration(exp)
	}
	// if jwt-secret given then use it
	if jwtSecret != "" {
		env.JwtSecret = jwtSecret
		SignedString = []byte(jwtSecret)
	}
	if url == "" {
		log.Fatalf("%s is required!", EnvDbUrl)
	}
	if dbName == "" {
		log.Fatalf("%s is required!", EnvDbName)
	}
	if collection == "" {
		log.Fatalf("%s is required!", EnvDbCollection)
	}
	env.DbUrl = url
	env.DbName = dbName
	env.DbUsername = username
	env.DbPassword = password
	env.DbCollection = collection
	E = env
	return env
}
