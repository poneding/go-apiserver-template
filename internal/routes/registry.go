package routes

import (
	"go-apiserver-template/docs"
	"go-apiserver-template/internal/global"
	v1 "go-apiserver-template/internal/routes/api/v1"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginswag "github.com/swaggo/gin-swagger"
)

// Register register all routers
func Register(root *gin.Engine) {
	root.Use(cors.Default())
	root.GET("/version", Version)
	root.GET(global.HealthzRelativePath, Healthz)

	if os.Getenv(global.RunModeEnvKey) == global.RunModeDev {
		if v := os.Getenv(global.ExternalHostEnvKey); len(v) > 0 {
			docs.SwaggerInfo.Host = v
		}
		root.GET("/swagger/*any", ginswag.WrapHandler(swagfiles.Handler))
	}

	apiv1 := root.Group("/api/v1")

	registerUserRouters(apiv1)
	registerUserOrderRouters(apiv1)
}

// registerUserRouters register user routers
// path: /api/v1/users
func registerUserRouters(base *gin.RouterGroup) {
	r := base.Group("/users")

	r.GET("", v1.ListUsers)
	r.GET("/:uid", v1.GetUser)
	r.POST("", v1.CreateUser)
	r.PUT("/:uid", v1.UpdateUser)
	r.DELETE("/:uid", v1.DeleteUser)
}

// registerUserOrderRouters register user order routers
// path: /api/v1/users/:uid/orders
func registerUserOrderRouters(base *gin.RouterGroup) {
	r := base.Group("/orders")

	r.GET("/:oid", v1.GetOrder)
}
