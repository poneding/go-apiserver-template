package routes

import (
	"go-apiserver-template/internal/global"
	"go-apiserver-template/internal/http"

	"github.com/gin-gonic/gin"
)

// Version returns the current version of the apisever
func Version(c *gin.Context) {
	http.Success(c, gin.H{
		"version": global.CurrentVersion,
	})
}
