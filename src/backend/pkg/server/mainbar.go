package server

import "github.com/gin-gonic/gin"

type MainServer struct {
}
type MainServerInterface interface {
	Get(c *gin.Context)
	Patch(c *gin.Context)
	Offer(c *gin.Context)
	Reject(c *gin.Context)
	NewSelect(c *gin.Context)
}
