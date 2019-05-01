package main

import (
	"fmt"
	"github.com/erkanzileli/go-auth"
	"github.com/erkanzileli/go-auth/controller"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()

	env, err := auth.GetEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Environment variables was taken.")

	// Initialize
	auth.InitializeTokenService(env)
	db := auth.Initialize(env.DbUrl, env.DbName,
		env.DbUsername, env.DbPassword, env.DbCollection)
	controller.Initialize(r, db)

	port := 8001

	if os.Getenv("GIN_MODE") == "release" {
		port = 80
	}

	// Run server
	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Print("Server is running.")

}
