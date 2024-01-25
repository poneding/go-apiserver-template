package routes

import (
	"go-apiserver-template/internal/api"
	"go-apiserver-template/internal/ctx"
	"go-apiserver-template/internal/global"

	"github.com/gin-gonic/gin"
)

// Version returns the current version of the apisever
func Version(c *ctx.Context) {
	api.Success(c, gin.H{
		"version": global.CurrentVersion,
	})
}
