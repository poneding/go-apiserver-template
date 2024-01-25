package routes

import (
	"net/http"

	"go-apiserver-template/internal/api"
	"go-apiserver-template/internal/ctx"
	"go-apiserver-template/internal/db"
)

// Healthz is a health check handler
func Healthz(c *ctx.Context) {
	if err := db.Ping(); err != nil {
		api.ErrorWithStatus(c, http.StatusInternalServerError, err)
		return
	}

	api.Success(c, "ok")
}
