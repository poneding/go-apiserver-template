package routes

import (
	stdhttp "net/http"

	"go-apiserver-template/internal/db"
	"go-apiserver-template/internal/http"

	"github.com/gin-gonic/gin"
)

// Healthz is a health check handler
func Healthz(c *gin.Context) {
	if err := db.Ping(); err != nil {
		http.ErrorWithStatus(c, stdhttp.StatusInternalServerError, err)
		return
	}

	http.Success(c, "ok")
}
