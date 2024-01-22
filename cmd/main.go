package main

import (
	_ "go-apiserver-template/docs"
	"go-apiserver-template/internal/server"
)

// @title Go Web Template API
// @description This is a sample server celler server.
// @version v1
// @host localhost:8080
func main() {
	apiserver := server.New()

	// http
	apiserver.Run(":8080")

	// https
	// apiserver.RunTLS(":8443", "cert.pem", "key.pem")
}
