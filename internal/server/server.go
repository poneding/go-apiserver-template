package server

import (
	"go-apiserver-template/internal/global"
	"go-apiserver-template/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server is the interface for http web server
// type Server interface {
// 	Run() error
// 	RunTLS(certFile, keyFile string)
// }

// serverOption is the option for http web server
type serverOption func(*server)

// WithPort set the port for http web server
func WithPort(port string) serverOption {
	return func(s *server) {
		s.Addr = port
	}
}

// New create a new http web server
func New(opts ...serverOption) *server {
	e := gin.New()
	e.Use(
		// builtin middlewares
		gin.LoggerWithWriter(gin.DefaultWriter, global.HealthzRelativePath), // ignore healthz log
		gin.Recovery(),
	)

	// Registry routers before run
	routes.Register(e)

	return &server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: e,
		},
	}
}

// server is the default implementation of Server interface
type server struct {
	*http.Server
}

// // Run run http web server
// func (s *server) Run() error {
// 	return s.Server.ListenAndServe()
// }

// // RunTLS run https web server
// func (s *server) RunTLS(certFile, keyFile string) {
// 	s.Server.ListenAndServeTLS(certFile, keyFile)
// }
