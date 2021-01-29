package main

import (
	"gofirebase/api"
	"gofirebase/config"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialize new gin engine (for server)
	r := gin.Default()

	//create/configure database instance
	db := config.SetupDB()

	//configure firebase
	firebaseAuth := config.SetupFirebase()

	//set db to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	//routes definition for finding and creating artists
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	//start server
	r.Run(":5000")
}
