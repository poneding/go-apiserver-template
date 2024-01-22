package v1

import (
	"go-apiserver-template/internal/http"
	"go-apiserver-template/internal/models"
	"go-apiserver-template/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// @Summary List users
// @Description get users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=[]models.User}
// @Failure 200 {object} http.Response
// @Router /api/v1/users [get]
func ListUsers(c *gin.Context) {
	users, err := services.NewUserService().ListUsers()
	if err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, users)
}

// @Summary Get user by id
// @Description get user by id
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "User ID"
// @Success 200 {object} http.Response{data=models.User}
// @Failure 200 {object} http.Response
// @Router /api/v1/users/{uid} [get]
func GetUser(c *gin.Context) {
	userID := cast.ToUint64(c.Param("uid"))
	u, err := services.NewUserService().GetUser(userID)
	if err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, u)
}

// @Summary Create user
// @Description create user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} http.Response{data=models.User}
// @Failure 200 {object} http.Response
// @Router /api/v1/users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		http.BadRequest(c)
		return
	}

	u, err := services.NewUserService().CreateUser(&user)
	if err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, u)
}

// @Summary Update user
// @Description update user
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} http.Response{data=models.User}
// @Failure 200 {object} http.Response
// @Router /api/v1/users/{uid} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		http.BadRequest(c)
		return
	}

	u, err := services.NewUserService().UpdateUser(&user)
	if err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, u)
}

// @Summary Delete user
// @Description delete user
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "User ID"
// @Success 200 {object} http.Response
// @Failure 200 {object} http.Response
// @Router /api/v1/users/{uid} [delete]
func DeleteUser(c *gin.Context) {
	userID := cast.ToUint64(c.Param("uid"))
	if err := services.NewUserService().DeleteUser(userID); err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, nil)
}
