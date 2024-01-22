package routes

import (
	"go-apiserver-template/internal/global"
	v1 "go-apiserver-template/internal/routes/api/v1"
	"os"

	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginswag "github.com/swaggo/gin-swagger"
)

// Register register all routers
func Register(root *gin.Engine) {
	root.GET("/version", Version)
	root.GET(global.HealthzRelativePath, Healthz)

	if os.Getenv(global.RunModeEnvKey) == global.RunModeDev {
		root.GET("/swagger/*any", ginswag.WrapHandler(swagfiles.Handler))
	}

	apiv1 := root.Group("/api/v1")

	registerUserRouters(apiv1)
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

	registerUserOrderRouters(r)
}

// registerUserOrderRouters register user order routers
// path: /api/v1/users/:uid/orders
func registerUserOrderRouters(base *gin.RouterGroup) {
	r := base.Group("/:uid/orders")

	r.GET("/:oid", v1.GetUserOrder)
}
