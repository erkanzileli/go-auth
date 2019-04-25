package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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
func GetEnv() (*Environment, error) {
	var e Environment
	var err error
	username := os.Getenv(EnvDbUsername)
	password := os.Getenv(EnvDbPassword)
	url := os.Getenv(EnvDbUrl)
	dbName := os.Getenv(EnvDbName)
	collection := os.Getenv(EnvDbCollection)
	jwtSecret := os.Getenv(EnvJwtSecret)
	jwtExpire := os.Getenv(EnvJwtExpire)

	if jwtExpire != "" {
		exp, err := strconv.Atoi(jwtExpire)
		if err != nil {
			log.Fatalf("%s must be int!", EnvJwtExpire)
			return nil, err
		}
		e.JwtExpire = time.Duration(exp)
	}
	// if jwt-secret given then use it
	if jwtSecret != "" {
		e.JwtSecret = jwtSecret
		SignedString = []byte(jwtSecret)
	}
	if url == "" {
		err = errors.New(fmt.Sprintf("%s is required!", EnvDbUrl))
	}
	if dbName == "" {
		err = errors.New(fmt.Sprintf("%s is required!", EnvDbName))
	}
	if collection == "" {
		err = errors.New(fmt.Sprintf("%s is required!", EnvDbCollection))
	}
	e.DbUrl = url
	e.DbName = dbName
	e.DbUsername = username
	e.DbPassword = password
	e.DbCollection = collection
	E = &e
	return &e, err
}
