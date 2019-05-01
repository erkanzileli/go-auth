package controller

import (
	"context"
	"fmt"
	"github.com/erkanzileli/go-auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// AuthData is represents the auth stack
type RequestModel struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login function is a controller for /login endpoint
func Login(c *gin.Context) {
	var authData RequestModel

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
		token, err := auth.CreateToken()
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
		"Error": "Email or password is wrong!",
	})

}

// FindUser represents the of there is a user if exist
func FindUser(email, password string) bool {
	var result RequestModel
	filter := bson.D{
		{"email", email},
		{"password", password}}
	err := Db.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	return true
}
