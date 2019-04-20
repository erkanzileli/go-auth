package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kalderasoft/go-auth"
)

var Db *auth.Database

// Initialize initializes controllers with router and db
func Initialize(r *gin.Engine, db *auth.Database) {

	// Routes
	r.POST("/login", Login)

	// Database Connection
	Db = db
}
