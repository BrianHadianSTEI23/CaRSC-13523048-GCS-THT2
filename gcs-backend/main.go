package main

import "github.com/gin-gonic/gin"

// init() : a function to setup the db by filling it with db schema needed or whatever
func init() {
	// bikin sesuai desc di atas
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":9090") // listens on 0.0.0.0:8080 by default
}
