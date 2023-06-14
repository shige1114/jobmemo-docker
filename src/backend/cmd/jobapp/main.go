package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	serv := gin.Default()
	serv.GET("/", func(c *gin.Context) {
		c.JSON(
			200, gin.H{
				"messege": "pong",
			})
	})

	fmt.Println("Hello, World!")
	var i string

	fmt.Println(i)

}