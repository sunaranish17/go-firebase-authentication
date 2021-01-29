package main

import "github.com/gin-gonic/gin"

func main() {
	//initialize new gin engine (for server)
	r := gin.Default()

	//create/configure database instance
	r.Run(":5000")
}
