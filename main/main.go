package main

import (
	"auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	r.POST("/login", Login)

	err := r.Run(":4000") // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	token, err := auth.CreateToken(email, password)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"Authorization": fmt.Sprintf("JWT %s", token),
	})
}
