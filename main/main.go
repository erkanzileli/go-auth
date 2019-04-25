package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kalderasoft/go-auth"
	"github.com/kalderasoft/go-auth/controller"
)

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()

	env := auth.GetEnv()
	log.Print("Environment variables was taken.")

	db := auth.Initialize(env.DbUrl, env.DbName,
		env.DbUsername, env.DbPassword, env.DbCollection)

	controller.Initialize(r, db)

	port := 8001

	if os.Getenv("GIN_MODE") == "release" {
		port = 80
	}

	// Run server
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Print("Server is running.")

}
