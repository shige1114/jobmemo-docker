package server

import "github.com/gin-gonic/gin"

type SideServer struct {
}

type SideServerInterface interface {
}

func (s *SideServer) Get(c *gin.Context)      {}
func (s *SideServer) Register(c *gin.Context) {}
func (s *SideServer) Patch(c *gin.Context)    {}
func (s *SideServer) Delete(c *gin.Context)   {}
