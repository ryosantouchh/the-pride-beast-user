package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
	// db  *gorm.DB
}

// func NewServerApp(db *gorm.DB) *Server {
func NewServerApp() *Server {
	ginApp := gin.Default()

	return &Server{
		app: ginApp,
		// db:  db,
	}
}

func (s *Server) Start(port string) {
	// register route here

	// ":3001"
	s.HttpListenPort(port)
}

func (s *Server) HttpListenPort(port string) {
	err := s.app.Run(port)
	if err != nil {
		panic("cannot start user service")
	}
}
