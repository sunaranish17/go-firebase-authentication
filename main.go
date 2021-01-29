package main

import (
	"fmt"
	"gofirebase/config"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialize new gin engine (for server)
	r := gin.Default()

	//create/configure database instance
	db := config.SetupDB()
	fmt.Println(db)
	r.Run(":5000")
}
