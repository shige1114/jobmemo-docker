// advanced-middleware.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.POST("/test", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/sidebar", func(c *gin.Context) {

	})
	r.POST("sidebar", func(ctx *gin.Context) {

	})
	r.PATCH("/sidebar", func(ctx *gin.Context) {})
	r.PATCH("/sidebar/:companies_id", func(ctx *gin.Context) {})

	r.GET("/main")
	r.POST("/main")
	r.PATCH("/main")

	r.Run(":8080")

}
