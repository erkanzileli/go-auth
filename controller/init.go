package controller

import (
	"github.com/erkanzileli/go-auth"
	"github.com/gin-gonic/gin"
)

// Db variable is represent the DB client for controllers
var Db *auth.Database

// Initialize initializes controllers with router and db
func Initialize(r *gin.Engine, db *auth.Database) {

	// Routes
	r.POST("/login", Login)

	// Temporary
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up",
		})
	})

	// Database Connection
	Db = db
}
