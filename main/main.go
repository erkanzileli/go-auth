package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalderasoft/go-auth"
	"github.com/kalderasoft/go-auth/controller"
	"log"
)

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()

	env := auth.GetEnv()
	log.Print("Environment variables was taken.")

	db := auth.Initialize(env.DbUrl, env.DbName,
		env.DbUsername, env.DbPassword, env.DbCollection)
	log.Print("Database connection is established.")

	controller.Initialize(r, db)

	// Run server
	err := r.Run(":8000")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Print("Server is running.")

}

