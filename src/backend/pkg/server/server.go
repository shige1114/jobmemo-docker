package server

import "github.com/gin-gonic/gin"

type Server struct {
	server *gin.Engine
	side   *SideServer
	main   *MainServer
}
type ServerInterface interface{}
