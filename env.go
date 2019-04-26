package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	JwtExpire    int
}

var E *Environment

// GetEnv func takes all needed environment variables fro OS
func GetEnv() (*Environment, error) {
	env := new(Environment)
	var err error
	username := os.Getenv(EnvDbUsername)
	password := os.Getenv(EnvDbPassword)
	url := os.Getenv(EnvDbUrl)
	dbName := os.Getenv(EnvDbName)
	collection := os.Getenv(EnvDbCollection)
	jwtSecret := os.Getenv(EnvJwtSecret)
	jwtExpire := os.Getenv(EnvJwtExpire)

	if jwtExpire != "" {
		exp, e := strconv.Atoi(jwtExpire)
		if exp <= 0 {
			err = errors.New(fmt.Sprintf("%s must greather than zero!", EnvJwtExpire))
		}
		if e != nil {
			err = errors.New(fmt.Sprintf("%s must be int!", EnvJwtExpire))
		}
		env.JwtExpire = exp
	}
	if jwtSecret != "" {
		env.JwtSecret = jwtSecret
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

	env.DbUrl = url
	env.DbName = dbName
	env.DbUsername = username
	env.DbPassword = password
	env.DbCollection = collection

	E = env
	return env, err
}
