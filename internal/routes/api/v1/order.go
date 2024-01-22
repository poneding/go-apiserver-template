package v1

import (
	"go-apiserver-template/internal/http"
	"go-apiserver-template/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// @Summary Get user order by id
// @Description Get user order by id
// @Tags User
// @Accept json
// @Produce json
// @Param uid path int true "User ID"
// @Param oid path int true "Order ID"
// @Success 200 {object} http.Response{data=models.Order}
// @Failure 200 {object} http.Response
// @Router /api/v1/users/{uid}/orders/{oid} [get]
func GetUserOrder(c *gin.Context) {
	uid := cast.ToUint64(c.Param("uid"))
	oid := cast.ToUint64(c.Param("oid"))

	o, err := services.NewOrderService().GetUserOrder(uid, oid)
	if err != nil {
		http.Error(c, err)
		return
	}

	http.Success(c, o)
}
