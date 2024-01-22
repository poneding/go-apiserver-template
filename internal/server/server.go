package server

import (
	"go-apiserver-template/internal/global"
	"go-apiserver-template/internal/routes"

	"github.com/gin-gonic/gin"
)

// Server is the interface for http web server
type Server interface {
	Run(addr ...string)
	RunTLS(addr, certFile, keyFile string)
}

// New create a new http web server
func New() Server {
	e := gin.New()
	e.Use(
		// builtin middlewares
		gin.LoggerWithWriter(gin.DefaultWriter, global.HealthzRelativePath), // ignore healthz log
		gin.Recovery(),
	)

	// Registry routers before run
	routes.Register(e)

	return &server{
		Engine: e,
	}
}

// server is the default implementation of Server interface
type server struct {
	*gin.Engine
}

// Run run http web server
func (s *server) Run(addr ...string) {
	s.Engine.Run(addr...)
}

// RunTLS run https web server
func (s *server) RunTLS(addr, certFile, keyFile string) {
	s.Engine.RunTLS(addr, certFile, keyFile)
}
