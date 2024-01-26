package server

import (
	"go-apiserver-template/internal/global"
	"go-apiserver-template/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerOption is the option for http web server
type ServerOption func(*Server)

// WithPort set the port for http web server
func WithPort(port string) ServerOption {
	return func(s *Server) {
		s.Addr = port
	}
}

// New create a new http web server
func New(opts ...ServerOption) *Server {
	e := gin.New()
	e.Use(
		// builtin middlewares
		gin.LoggerWithWriter(gin.DefaultWriter, global.HealthzRelativePath), // ignore healthz log
		gin.Recovery(),
	)

	// disable gin bind validation
	// gin.DisableBindValidation()
	// disable gin color log
	// gin.DisableConsoleColor()

	// custom log writer
	// f, _ := os.Create("server.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Registry routers before run
	routes.Register(e)

	return &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: e,
		},
	}
}

// Server is the default implementation of Server interface
type Server struct {
	*http.Server
}
