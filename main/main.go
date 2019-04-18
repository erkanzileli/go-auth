package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kalderasoft/go-auth"
)

func main() {
	// // Force log's color
	// gin.ForceConsoleColor()

	// // Creates a gin router with default middleware:
	// r := gin.Default()

	// err := r.Run(":4000") // listen and serve on 0.0.0.0:8080
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	auth.InitDatabase()

}

// Login function is a controller for /login endpoint
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
