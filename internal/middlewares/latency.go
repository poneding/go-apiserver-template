package middlewares

import (
	"go-apiserver-template/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
)

func Latency() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Infof("%s %s %s", c.Request.Method, c.Request.URL.Path, latency)
	}
}
