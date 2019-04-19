package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kalderasoft/go-auth"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type AuthData struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var DB *auth.Database

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()

	env := auth.GetEnv()
	DB = auth.InitDatabase(env.DB_URL, env.DB_NAME,
		env.DB_USERNAME, env.DB_PASSWORD, env.DB_COLLECTION)

	// Routes
	r.POST("/login", Login)

	//err := autotls.Run(r, "example1.com", "example2.com") // listen and serve on 0.0.0.0:8080

	// Run server
	err := r.Run(":8000")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

}

// Login function is a controller for /login endpoint
func Login(c *gin.Context) {
	var authData AuthData

	err := c.BindJSON(&authData)
	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	userFind := FindUser(authData.Email, authData.Password)
	if userFind {
		token, err := auth.CreateToken(authData.Email, authData.Password)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Authorization": fmt.Sprintf("JWT %s", token),
		})
		return
	}
	c.JSON(200, gin.H{
		"Error": "That is have a mistake on Username or Password",
	})

}

func FindUser(email, password string) bool {
	var result AuthData
	filter := bson.D{
		{"email", email},
		{"password", password}}
	err := DB.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	return true
}
