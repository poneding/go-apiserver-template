package main

import (
	"context"
	"go-apiserver-template/internal/server"
	"go-apiserver-template/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Go APIServer Template
// @description This is a sample server celler server.
// @version v1
// @host localhost:8080
func main() {
	// http
	srv := server.New(server.WithPort(":8080"))

	go func() {
		// graceful shutdown
		log.Info("Start Server ...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// https
	// srv := server.New(server.WithPort(":8080"))
	//
	// go func() {
	// 	// graceful shutdown
	// 	if err :=
	// 		srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// wait for interrupt signal to gracefully shutdown the server
	// with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
