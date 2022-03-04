package utility

import "github.com/gin-gonic/gin"

type MyServer struct {
	server *gin.Engine
}

func (s *MyServer) AddRoute(path string, routeGroupFunc func(*gin.RouterGroup)) {
	routeGroupFunc(s.server.Group(path))
}

func (s *MyServer) Run() {
	s.server.Run(":8080")
}

func NewMyServer(defaultServer bool) *MyServer {
	if defaultServer {
		return &MyServer{server: gin.Default()}
	} else {
		return &MyServer{server: gin.New()}
	}
}
