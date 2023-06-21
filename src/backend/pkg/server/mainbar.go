package server

import (
	"github.com/backend"
	"github.com/gin-gonic/gin"
)

type MainServer struct {
}
type MainServerInterface interface {
	Route()
	Get(c *gin.Context) backend.Main
	Patch(c *gin.Context)
	Offer(c *gin.Context)
	Reject(c *gin.Context)
	NewSelect(c *gin.Context)
}
