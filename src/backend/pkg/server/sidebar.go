package server

import "github.com/gin-gonic/gin"

type SideServer struct {
	engine *gin.Engine
}

type SideServerInterface interface {
	Route()
	Get(c *gin.Context)
	Register(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
}

func (s *SideServer) Route()                  {}
func (s *SideServer) Get(c *gin.Context)      {}
func (s *SideServer) Register(c *gin.Context) {}
func (s *SideServer) Patch(c *gin.Context)    {}
func (s *SideServer) Delete(c *gin.Context)   {}
